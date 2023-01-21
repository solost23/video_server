package middlewares

import (
	"errors"
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"video_server/global"
	"video_server/pkg/response"
	"video_server/pkg/utils"
)

func AuthCheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := utils.GetUser(c).Role
		mysqlConf := global.ServerConfig.MysqlConfig
		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", mysqlConf.User, mysqlConf.Password, mysqlConf.Addr, mysqlConf.DB)
		a, err := gormadapter.NewAdapter("mysql", dsn, true)
		if err != nil {
			response.Error(c, 2001, err)
			return
		}
		e, err := casbin.NewEnforcer("./configs/rbac_model.conf", a)
		if err != nil {
			response.Error(c, 2001, err)
			return
		}
		err = e.LoadPolicy()
		if err != nil {
			response.Error(c, 2001, err)
			return
		}
		ok, err := e.Enforce(role, c.Request.URL.Path, c.Request.Method)
		if err != nil {
			response.Error(c, 2001, err)
			return
		}
		if !ok {
			response.Error(c, 2001, errors.New("权限认证不通过"))
			return
		}
		c.Next()
		return
	}
}
