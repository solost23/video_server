package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"mime/multipart"
	"strconv"
	"strings"
	"time"
	"video_server/forms"
	"video_server/global"
	"video_server/pkg/cache"
	"video_server/pkg/constants"
	"video_server/pkg/middlewares"
	"video_server/pkg/models"
	"video_server/pkg/utils"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (s *Service) Register(c *gin.Context, params *forms.RegisterForm) (err error) {
	// base logic: 校验当前用户是否存在，若不存在则新建
	db := global.DB

	query := []string{"user_name = ?"}
	args := []interface{}{params.UserName}
	_, err = (&models.User{}).WhereOne(db, strings.Join(query, " AND "), args...)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if err == nil {
		return errors.New("用户已存在")
	}
	user := &models.User{
		UserName:     params.UserName,
		Password:     utils.NewMd5(params.Password, global.ServerConfig.Md5Config.Secret),
		Nickname:     params.Nickname,
		Role:         params.Role,
		Avatar:       params.Avatar,
		Introduce:    params.Introduce,
		FansCount:    0,
		CommentCount: 0,
	}
	err = user.Insert(db)
	if err != nil {
		return err
	}

	z := &Zinc{Username: global.ServerConfig.ZincConfig.Username, Password: global.ServerConfig.ZincConfig.Password}
	err = z.InsertDocument(c, constants.ZINCINDEXUSER, strconv.Itoa(int(user.ID)), map[string]interface{}{
		"username":  user.UserName,
		"nickname":  user.Nickname,
		"role":      user.Role,
		"introduce": user.Introduce,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Login(c *gin.Context, params *forms.LoginForm) (response *forms.LoginResponse, err error) {
	db := global.DB

	user, err := (&models.User{}).WhereOne(db, "user_name = ?", params.UserName)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New(fmt.Sprintf("用户%s不存在", params.UserName))
	}
	if params.UserName != user.UserName || utils.NewMd5(params.Password, global.ServerConfig.Md5Config.Secret) != user.Password {
		return nil, errors.New("用户名或密码错误")
	}
	// 区分两种设备 分别是 web 和 mobile
	var redisPrefix string
	if params.Device == "web" {
		redisPrefix = constants.WebRedisPrefix
	} else {
		redisPrefix = constants.MobileRedisPrefix
	}

	j := middlewares.NewJWT()
	claims := middlewares.CustomClaims{
		UserId: user.ID,
		Device: redisPrefix,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + int64(global.ServerConfig.JWTConfig.Duration),
			Issuer:    "video_server",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		return nil, err
	}

	userJson, _ := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	rdb, err := cache.RedisConnFactory(15)
	if err != nil {
		return nil, err
	}

	key := redisPrefix + strconv.Itoa(int(user.ID))
	oldToken, err := rdb.Get(c, key).Result()

	rdb.Del(c, constants.RedisPrefix+oldToken)
	rdb.Set(c, key, token, time.Duration(global.ServerConfig.JWTConfig.Duration)*time.Second)
	rdb.Set(c, constants.RedisPrefix+token, userJson, time.Duration(global.ServerConfig.JWTConfig.Duration)*time.Second)

	if err = user.Updates(db, map[string]interface{}{"last_login_time": time.Now()}, "id = ?", user.ID); err != nil {
		return nil, err
	}

	// 封装数据返回
	response = &forms.LoginResponse{
		IsFirstLogin: 2,
		User:         *user,
		Token:        token,
	}
	return response, err
}

func (s *Service) Logout(c *gin.Context, params *forms.LogoutForm) (err error) {
	user := c.Value("user").(*models.User)
	rdb, err := cache.RedisConnFactory(15)
	if err != nil {
		return err
	}
	var redisPrefix string
	if params.Device == "web" {
		redisPrefix = constants.WebRedisPrefix
	} else {
		redisPrefix = constants.MobileRedisPrefix
	}
	key := redisPrefix + strconv.Itoa(int(user.ID))
	token, err := rdb.Get(c, key).Result()
	if err != nil {
		return err
	}
	rdb.Del(c, constants.RedisPrefix+token)
	return nil
}

func (s *Service) ListUser(c *gin.Context, params *forms.ListForm) (response *forms.ListResponse, err error) {
	db := global.DB

	query := make([]string, 0, 3)
	args := make([]interface{}, 0, 3)

	if params.ID > 0 {
		query = append(query, "id = ?")
		args = append(args, params.ID)
	}
	if params.UserName != "" {
		query = append(query, "user_name LIKE ?")
		args = append(args, models.LikeFilter(params.UserName))
	}
	if params.Role != "" {
		query = append(query, "role LIKE ?")
		args = append(args, models.LikeFilter(params.Role))
	}
	users, total, err := (&models.User{}).PageListOrder(db, "", &models.ListPageInput{
		Page: params.Page,
		Size: params.Size,
	}, strings.Join(query, " AND "), args...)
	// 封装数据返回
	records := make([]forms.ListRecord, 0, len(users))
	for _, user := range users {
		records = append(records, forms.ListRecord{
			ID:           user.ID,
			UserName:     user.UserName,
			Nickname:     user.Nickname,
			Role:         user.Role,
			Avatar:       user.Avatar,
			Introduce:    user.Introduce,
			FansCount:    user.FansCount,
			CommentCount: user.CommentCount,
			CreateTime:   user.CreatedAt.Format(constants.TimeFormat),
			UpdateTime:   user.CreatedAt.Format(constants.TimeFormat),
		})
	}
	response = &forms.ListResponse{
		List: records,
		PageList: &utils.PageList{
			Size:    params.Size,
			Pages:   int64(math.Ceil(float64(total) / float64(params.Size))),
			Total:   total,
			Current: params.Page,
		},
	}
	return response, nil
}

func (s *Service) DeleteUser(c *gin.Context, id uint) (err error) {
	// 查看用户是否存在，若存在，则删除
	db := global.DB

	query := []string{"id = ?"}
	args := []interface{}{id}
	_, err = (&models.User{}).WhereOne(db, strings.Join(query, " AND "), args...)
	if err != nil {
		return err
	}

	tx := db.Begin()
	// 删除用户下的分类
	query = []string{"user_id IN ?"}
	err = (&models.Category{}).Delete(tx, strings.Join(query, " AND "), args...)
	if err != nil {
		tx.Rollback()
		return err
	}
	// 删除用户下的视频
	err = (&models.Video{}).Delete(tx, strings.Join(query, " AND "), args...)
	if err != nil {
		tx.Rollback()
		return err
	}
	// 删除用户的所有评论
	err = (&models.Comment{}).Delete(tx, strings.Join(query, " AND "), args...)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	// 从全局索引中删除用户记录
	z := &Zinc{Username: global.ServerConfig.ZincConfig.Username, Password: global.ServerConfig.ZincConfig.Password}
	err = z.DeleteDocument(c, constants.ZINCINDEXUSER, strconv.Itoa(int(id)))
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateUser(c *gin.Context, id uint, params *forms.UserUpdateForm) (err error) {
	db := global.DB

	query := []string{"id = ?"}
	args := []interface{}{id}
	// base logic: 检查用户是否存在，若存在，则删除
	_, err = (&models.User{}).WhereOne(db, strings.Join(query, " AND "), args...)
	if err != nil {
		return err
	}
	value := map[string]interface{}{
		"user_name": params.UserName,
		"password":  utils.NewMd5(params.Password, global.ServerConfig.Md5Config.Secret),
		"nickname":  params.Nickname,
		"avatar":    params.Avatar,
		"introduce": params.Introduce,
	}
	err = (&models.User{}).Updates(db, value, strings.Join(query, " AND "), args...)
	if err != nil {
		return err
	}
	user, err := (&models.User{}).WhereOne(db, strings.Join(query, " AND "), args...)
	if err != nil {
		return err
	}
	z := &Zinc{Username: global.ServerConfig.ZincConfig.Username, Password: global.ServerConfig.ZincConfig.Password}
	err = z.DeleteDocument(c, constants.ZINCINDEXUSER, strconv.Itoa(int(user.ID)))
	if err != nil {
		return err
	}
	err = z.InsertDocument(c, constants.ZINCINDEXUSER, strconv.Itoa(int(user.ID)), map[string]interface{}{
		"username":  user.UserName,
		"nickname":  user.Nickname,
		"role":      user.Role,
		"introduce": user.Introduce,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Detail(c *gin.Context, id uint) (response *forms.ListRecord, err error) {
	db := global.DB

	query := []string{"id = ?"}
	args := []interface{}{id}
	user, err := (&models.User{}).WhereOne(db, strings.Join(query, " AND "), args...)
	if err != nil {
		return nil, err
	}
	// 封装数据，返回
	response = &forms.ListRecord{
		ID:           user.ID,
		UserName:     user.UserName,
		Nickname:     user.Nickname,
		Role:         user.Role,
		Avatar:       user.Avatar,
		Introduce:    user.Introduce,
		FansCount:    user.FansCount,
		CommentCount: user.CommentCount,
		CreateTime:   user.CreatedAt.Format(constants.TimeFormat),
		UpdateTime:   user.UpdatedAt.Format(constants.TimeFormat),
	}
	return response, nil
}

func (s *Service) UploadAvatar(c *gin.Context, file *multipart.FileHeader) (result string, err error) {
	user := &models.User{}
	result, err = UploadImg(user, "avatar", file)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (s *Service) SearchUser(c *gin.Context, params *forms.SearchForm) (*forms.ListResponse, error) {
	db := global.DB

	z := &Zinc{Username: global.ServerConfig.ZincConfig.Username, Password: global.ServerConfig.ZincConfig.Password}
	from := int32((params.Page - 1) * params.Size)
	size := from + int32(params.Size) - 1
	searchResults, total, err := z.SearchDocument(c, constants.ZINCINDEXUSER, params.Keyword, from, size)
	if err != nil {
		return nil, err
	}
	userIds := make([]uint, 0, len(searchResults))
	for _, searchResult := range searchResults {
		id, _ := strconv.Atoi(*searchResult.Id)
		userIds = append(userIds, uint(id))
	}
	users, err := (&models.User{}).WhereAll(db, "id IN ?", userIds)
	if err != nil {
		return nil, err
	}
	userIdToUserInfoMaps := make(map[uint]struct {
		Avatar       string
		FansCount    int64
		CommentCount int64
		CreateTime   time.Time
		UpdateTime   time.Time
	}, len(users))
	for _, user := range users {
		userIdToUserInfoMaps[user.ID] = struct {
			Avatar       string
			FansCount    int64
			CommentCount int64
			CreateTime   time.Time
			UpdateTime   time.Time
		}{Avatar: user.Avatar, FansCount: user.FansCount, CommentCount: user.CommentCount, CreateTime: user.CreatedAt, UpdateTime: user.UpdatedAt}
	}
	// 封装数据并返回
	records := make([]forms.ListRecord, 0, len(searchResults))
	for _, searchResult := range searchResults {
		id, _ := strconv.Atoi(*searchResult.Id)
		records = append(records, forms.ListRecord{
			ID:           uint(id),
			UserName:     searchResult.Source["username"].(string),
			Nickname:     searchResult.Source["nickname"].(string),
			Role:         searchResult.Source["role"].(string),
			Avatar:       userIdToUserInfoMaps[uint(id)].Avatar,
			Introduce:    searchResult.Source["introduce"].(string),
			FansCount:    userIdToUserInfoMaps[uint(id)].FansCount,
			CommentCount: userIdToUserInfoMaps[uint(id)].CommentCount,
			CreateTime:   userIdToUserInfoMaps[uint(id)].CreateTime.Format(constants.TimeFormat),
			UpdateTime:   userIdToUserInfoMaps[uint(id)].UpdateTime.Format(constants.TimeFormat),
		})
	}

	result := &forms.ListResponse{
		List: records,
		PageList: &utils.PageList{
			Size:    params.Size,
			Pages:   int64(math.Ceil(float64(total) / float64(size))),
			Total:   total,
			Current: params.Page,
		},
	}

	return result, nil
}
