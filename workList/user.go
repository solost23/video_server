package workList

import (
	"errors"
	"math"
	"strings"
	"video_server/forms"
	"video_server/pkg/middleware"
	"video_server/pkg/models"
	"video_server/pkg/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService struct {
	WorkList
}

func (w *UserService) Register(c *gin.Context, params *forms.RegisterForm) (response *forms.RegisterResponse, err error) {
	// base logic: 校验当前用户是否存在，若不存在则新建
	query := []string{"user_name = ?"}
	args := []interface{}{params.UserName}
	_, err = (&models.User{}).WhereOne(w.GetMysqlConn(), strings.Join(query, " AND "), args...)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if err == nil {
		return nil, errors.New("用户已存在")
	}
	err = (&models.User{
		UserName:     params.UserName,
		Password:     utils.NewMd5(params.Password, models.SECRET),
		Nickname:     params.Nickname,
		Role:         params.Role,
		Avatar:       params.Avatar,
		Introduce:    params.Introduce,
		FansCount:    0,
		CommentCount: 0,
	}).Insert(w.GetMysqlConn())
	if err != nil {
		return nil, err
	}
	return response, err
}

func (w *UserService) Login(c *gin.Context, params *forms.LoginForm) (response *forms.LoginResponse, err error) {
	// base logic: 查看有无用户，若有进行登录校验
	query := []string{"user_name = ?"}
	args := []interface{}{params.UserName}
	user, err := (&models.User{}).WhereOne(w.GetMysqlConn(), strings.Join(query, " AND "), args...)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if params.UserName != user.UserName || utils.NewMd5(params.Password, models.SECRET) != user.Password {
		return nil, errors.New("userName or password err")
	}
	tokenStr, err := middleware.CreateToken(user.UserName, user.Role)
	if err != nil {
		return nil, err
	}
	return &forms.LoginResponse{TokenStr: tokenStr}, err
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
			CreateTime:   user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdateTime:   user.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	response = &forms.ListResponse{
		List: records,
		PageList: &models.PageList{
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
		"password":  utils.NewMd5(params.Password, models.SECRET),
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
		CreateTime:   user.CreatedAt.Format(models.TimeFormat),
		UpdateTime:   user.UpdatedAt.Format(models.TimeFormat),
	}
	return response, nil
}
