package middleware

import (
	"net/http"

	"github.com/casbin/casbin"
	xormadapter "github.com/casbin/xorm-adapter"
	"github.com/gin-gonic/gin"
)

func AuthCheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, ok := c.Get("user")
		if !ok {
			c.Abort()
			return
		}
		role := claims.(*Claims).User.Role
		// 写入配置文件
		casbinDsn := "root:123@tcp(localhost:3306)/"
		a := xormadapter.NewAdapter("mysql", casbinDsn)
		e := casbin.NewEnforcer("../config/rbac_model.conf", a)
		if err := e.LoadPolicy(); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			c.Abort()
			return
		}
		ok = e.Enforce(role, c.Request.URL.Path, c.Request.Method)
		if !ok {
			c.JSON(http.StatusInternalServerError, "AuthCheckRole Failed")
			c.Abort()
			return
		}
		c.Next()
		return
	}
}
