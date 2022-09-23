package workList

import (
	"encoding/json"
	"errors"
	"math"
	"mime/multipart"
	"strconv"
	"strings"
	"time"
	"video_server/config"
	"video_server/forms"
	"video_server/pkg/cache"
	"video_server/pkg/constants"
	"video_server/pkg/middleware"
	"video_server/pkg/models"
	"video_server/pkg/utils"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService struct {
	WorkList
}

func (w *UserService) Register(c *gin.Context, params *forms.RegisterForm) (err error) {
	// base logic: 校验当前用户是否存在，若不存在则新建
	query := []string{"user_name = ?"}
	args := []interface{}{params.UserName}
	_, err = (&models.User{}).WhereOne(w.GetMysqlConn(), strings.Join(query, " AND "), args...)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if err == nil {
		return errors.New("用户已存在")
	}
	err = (&models.User{
		UserName:     params.UserName,
		Password:     utils.NewMd5(params.Password, constants.SECRET),
		Nickname:     params.Nickname,
		Role:         params.Role,
		Avatar:       params.Avatar,
		Introduce:    params.Introduce,
		FansCount:    0,
		CommentCount: 0,
	}).Insert(w.GetMysqlConn())
	if err != nil {
		return err
	}
	return nil
}

func (w *UserService) Login(c *gin.Context, params *forms.LoginForm) (response *forms.LoginResponse, err error) {
	user, err := (&models.User{}).WhereOne(w.GetMysqlConn(), "user_name = ?", params.UserName)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if params.UserName != user.UserName || utils.NewMd5(params.Password, constants.SECRET) != user.Password {
		return nil, errors.New("userName or password err")
	}
	// 区分两种设备 分别是 web 和 mobile
	var redisPrefix string
	if params.Device == "web" {
		redisPrefix = constants.WebRedisPrefix
	} else {
		redisPrefix = constants.MobileRedisPrefix
	}

	j := middleware.NewJWT()
	claims := middleware.CustomClaims{
		UserId: user.ID,
		Device: redisPrefix,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + config.NewJWTConfig().Duration,
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
	rdb.Set(c, key, token, time.Duration(config.NewJWTConfig().Duration)*time.Second)
	rdb.Set(c, constants.RedisPrefix+token, userJson, time.Duration(config.NewJWTConfig().Duration)*time.Second)

	if err = user.Updates(w.GetMysqlConn(), map[string]interface{}{"last_login_time": time.Now()}, "id = ?", user.ID); err != nil {
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

func (w *UserService) Logout(c *gin.Context, params *forms.LogoutForm) (err error) {
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

func (w *UserService) List(c *gin.Context, params *forms.ListForm) (response *forms.ListResponse, err error) {
	db := w.GetMysqlConn()

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

func (w *UserService) Delete(c *gin.Context, id uint) (err error) {
	// 查看用户是否存在，若存在，则删除
	db := w.GetMysqlConn()

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
	return nil
}

func (w *UserService) Update(c *gin.Context, id uint, params *forms.UserUpdateForm) (err error) {
	db := w.GetMysqlConn()

	query := []string{"id = ?"}
	args := []interface{}{id}
	// base logic: 检查用户是否存在，若存在，则删除
	_, err = (&models.User{}).WhereOne(db, strings.Join(query, " AND "), args...)
	if err != nil {
		return
	}
	value := map[string]interface{}{
		"user_name": params.UserName,
		"password":  utils.NewMd5(params.Password, constants.SECRET),
		"nickname":  params.Nickname,
		"avatar":    params.Avatar,
		"introduce": params.Introduce,
	}
	err = (&models.User{}).Updates(db, value, strings.Join(query, " AND "), args...)
	if err != nil {
		return err
	}
	return nil
}

func (w *UserService) Detail(c *gin.Context, id uint) (response *forms.ListRecord, err error) {
	db := w.GetMysqlConn()

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

func (w *UserService) UploadAvatar(c *gin.Context, file *multipart.FileHeader) (result string, err error) {
	user := &models.User{}
	result, err = UploadImg(user, "avatar", file)
	if err != nil {
		return "", err
	}

	return result, nil
}
