package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
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

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (s *Service) Register(c *gin.Context, params *forms.RegisterForm) (err error) {
	// base logic: 校验当前用户是否存在，若不存在则新建
	db := global.DB

	query := []string{"username = ?"}
	args := []interface{}{params.Username}
	sqlUser, err := (&models.User{}).WhereOne(db, strings.Join(query, " AND "), args...)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if sqlUser != nil && sqlUser.ID > 0 {
		return errors.New("用户已存在")
	}
	user := &models.User{
		Username:     *params.Username,
		Password:     utils.NewMd5(*params.Password, global.ServerConfig.Md5Config.Secret),
		Nickname:     *params.Nickname,
		Role:         *params.Role,
		Avatar:       utils.TrimDomainPrefix(*params.Avatar),
		Introduce:    *params.Introduce,
		FansCount:    0,
		CommentCount: 0,
	}

	if err := models.GInsert(db, user); err != nil {
		return err
	}

	z := NewZinc()
	err = z.InsertDocument(c, constants.ZINCINDEXUSER, strconv.Itoa(int(user.ID)), map[string]interface{}{
		"username":  user.Username,
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

	sqlUser, err := models.GWhereFirstSelect(db, &models.User{}, "*", "user_name = ?", *params.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if sqlUser == nil {
		return nil, errors.New(fmt.Sprintf("用户%s不存在", *params.Username))
	}
	if *params.Username != sqlUser.Username || utils.NewMd5(*params.Password, global.ServerConfig.Md5Config.Secret) != sqlUser.Password {
		return nil, errors.New("用户名或密码错误")
	}
	// 区分两种设备 分别是 web 和 mobile
	var redisPrefix string
	if *params.Device == "web" {
		redisPrefix = constants.WebRedisPrefix
	} else {
		redisPrefix = constants.MobileRedisPrefix
	}

	j := middlewares.NewJWT()
	claims := middlewares.CustomClaims{
		UserId: sqlUser.ID,
		Device: redisPrefix,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + int64(global.ServerConfig.JWTConfig.Duration),
			Issuer:    "video",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		return nil, err
	}

	userJson, _ := json.Marshal(sqlUser)
	if err != nil {
		return nil, err
	}

	rdb, err := cache.RedisConnFactory(15)
	if err != nil {
		return nil, err
	}

	key := redisPrefix + strconv.Itoa(int(sqlUser.ID))
	oldToken, err := rdb.Get(c, key).Result()

	rdb.Del(c, constants.RedisPrefix+oldToken)
	rdb.Set(c, key, token, time.Duration(global.ServerConfig.JWTConfig.Duration)*time.Second)
	rdb.Set(c, constants.RedisPrefix+token, userJson, time.Duration(global.ServerConfig.JWTConfig.Duration)*time.Second)

	err = models.GUpdateColumn(db, &models.User{}, "last_login_time", time.Now(), "id = ?", sqlUser.ID)
	if err != nil {
		return nil, err
	}

	sqlUser.Avatar = utils.FulfillImageOSSPrefix(sqlUser.Avatar)
	// 封装数据返回
	response = &forms.LoginResponse{
		User:  sqlUser,
		Token: &token,
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
	if *params.Device == "web" {
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

	if *params.ID > 0 {
		query = append(query, "id = ?")
		args = append(args, params.ID)
	}
	if *params.Username != "" {
		query = append(query, "user_name LIKE ?")
		args = append(args, models.LikeFilter(*params.Username))
	}
	if params.Role != nil {
		query = append(query, "role = ?")
		args = append(args, params.Role)
	}
	sqlUsers, total, err := (&models.User{}).PageListOrder(db, "", &models.ListPageInput{
		Page: params.Page,
		Size: params.Size,
	}, strings.Join(query, " AND "), args...)
	// 封装数据返回
	records := make([]*forms.ListRecord, 0, len(sqlUsers))
	for _, sqlUser := range sqlUsers {
		id := sqlUser.ID
		username := sqlUser.Username
		nickname := sqlUser.Nickname
		role := sqlUser.Role
		avatar := utils.FulfillImageOSSPrefix(sqlUser.Avatar)
		introduce := sqlUser.Introduce
		fansCount := sqlUser.FansCount
		commentCount := sqlUser.CommentCount
		createdAt := sqlUser.CreatedAt.Format(constants.DateTime)
		updatedAt := sqlUser.UpdatedAt.Format(constants.DateTime)
		records = append(records, &forms.ListRecord{
			ID:           &id,
			Username:     &username,
			Nickname:     &nickname,
			Role:         &role,
			Avatar:       &avatar,
			Introduce:    &introduce,
			FansCount:    &fansCount,
			CommentCount: &commentCount,
			CreatedAt:    &createdAt,
			UpdatedAt:    &updatedAt,
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
	z := NewZinc()
	err = z.DeleteDocument(c, constants.ZINCINDEXUSER, strconv.Itoa(int(id)))
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateUser(c *gin.Context, id uint, params *forms.UserUpdateForm) (err error) {
	db := global.DB

	sqlUser, err := models.GWhereFirstSelect(db, &models.User{}, "id", "id = ?", id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if sqlUser == nil {
		return errors.New("该用户不存在")
	}
	values := map[string]interface{}{
		"user_name": *params.Username,
		"password":  utils.NewMd5(*params.Password, global.ServerConfig.Md5Config.Secret),
		"nickname":  params.Nickname,
		"avatar":    utils.TrimDomainPrefix(*params.Avatar),
		"introduce": params.Introduce,
	}
	_, err = models.GUpdatesWhere(db, &models.User{}, values, "id = ?", id)
	if err != nil {
		return err
	}
	sqlUser, err = models.GWhereFirstSelect(db, &models.User{}, "*", "id = ?", id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if sqlUser == nil {
		return errors.New("该用户不存在")
	}

	z := NewZinc()
	err = z.DeleteDocument(c, constants.ZINCINDEXUSER, strconv.Itoa(int(sqlUser.ID)))
	if err != nil {
		return err
	}
	err = z.InsertDocument(c, constants.ZINCINDEXUSER, strconv.Itoa(int(sqlUser.ID)), map[string]interface{}{
		"username":  sqlUser.Username,
		"nickname":  sqlUser.Nickname,
		"role":      sqlUser.Role,
		"introduce": sqlUser.Introduce,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Detail(c *gin.Context, id uint) (response *forms.ListRecord, err error) {
	db := global.DB

	sqlUser, err := models.GWhereFirstSelect(db, &models.User{}, "*", "id = ?", id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if sqlUser == nil {
		return nil, errors.New("该用户不存在")
	}

	// 封装数据，返回
	avatar := utils.FulfillImageOSSPrefix(sqlUser.Avatar)
	createdAt := sqlUser.CreatedAt.Format(constants.DateTime)
	updatedAt := sqlUser.UpdatedAt.Format(constants.DateTime)
	response = &forms.ListRecord{
		ID:           &sqlUser.ID,
		Username:     &sqlUser.Username,
		Nickname:     &sqlUser.Nickname,
		Role:         &sqlUser.Role,
		Avatar:       &avatar,
		Introduce:    &sqlUser.Introduce,
		FansCount:    &sqlUser.FansCount,
		CommentCount: &sqlUser.CommentCount,
		CreatedAt:    &createdAt,
		UpdatedAt:    &updatedAt,
	}
	return response, nil
}

func (s *Service) UploadAvatar(c *gin.Context, file *multipart.FileHeader) (result string, err error) {
	folder := "video.server.users.avatar"

	url, err := UploadImg(0, folder, file.Filename, file, "image")
	if err != nil {
		return "", err
	}
	return utils.FulfillImageOSSPrefix(utils.TrimDomainPrefix(url)), nil
}

func (s *Service) SearchUser(c *gin.Context, params *forms.SearchForm) (*forms.ListResponse, error) {
	db := global.DB

	z := NewZinc()
	from := int32((params.Page - 1) * params.Size)
	size := from + int32(params.Size) - 1
	searchResults, total, err := z.SearchDocument(c, constants.ZINCINDEXUSER, *params.Keyword, from, size)
	if err != nil {
		return nil, err
	}
	userIds := make([]uint, 0, len(searchResults))
	for _, searchResult := range searchResults {
		id, _ := strconv.Atoi(*searchResult.Id)
		userIds = append(userIds, uint(id))
	}
	sqlUsers, err := models.GWhereAllSelectOrder(db, &models.User{}, "*", "id DESC", "id IN ?", userIds)
	if err != nil {
		return nil, err
	}

	records := make([]*forms.ListRecord, 0, len(searchResults))
	for _, sqlUser := range sqlUsers {
		id := sqlUser.ID
		username := sqlUser.Username
		nickname := sqlUser.Nickname
		role := sqlUser.Role
		avatar := utils.FulfillImageOSSPrefix(sqlUser.Avatar)
		introduce := sqlUser.Introduce
		fansCount := sqlUser.FansCount
		commentCount := sqlUser.CommentCount
		createdAt := sqlUser.CreatedAt.Format(constants.DateTime)
		updatedAt := sqlUser.UpdatedAt.Format(constants.DateTime)

		records = append(records, &forms.ListRecord{
			ID:           &id,
			Username:     &username,
			Nickname:     &nickname,
			Role:         &role,
			Avatar:       &avatar,
			Introduce:    &introduce,
			FansCount:    &fansCount,
			CommentCount: &commentCount,
			CreatedAt:    &createdAt,
			UpdatedAt:    &updatedAt,
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
