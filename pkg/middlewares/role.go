package middlewares

import (
	"errors"
	"video_server/pkg/models"
	"video_server/pkg/response"

	"github.com/casbin/casbin"
	xormadapter "github.com/casbin/xorm-adapter"
	"github.com/gin-gonic/gin"
)

func AuthCheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.Value("user").(*models.User)
		role := user.Role
		// 写入配置文件
		casbinDsn := "root:123@tcp(localhost:3306)/"
		a := xormadapter.NewAdapter("mysql", casbinDsn)
		e := casbin.NewEnforcer("../config/rbac_model.conf", a)
		if err := e.LoadPolicy(); err != nil {
			response.Error(c, 2001, err)
			return
		}
		ok := e.Enforce(role, c.Request.URL.Path, c.Request.Method)
		if !ok {
			response.Error(c, 2001, errors.New("权限认证错误"))
			return
		}
		c.Next()
		return
	}
}
