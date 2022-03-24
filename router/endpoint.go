package router

import (
	"github.com/gin-gonic/gin"

	"video_server/pkg/middleware"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	group := router.Group("")
	initNoAuthRouter(group)
	// 注意 role 需要再思考一下，不一定要放在这里
	//group.Use(jwt.JWTAuth, role.CheckRole)
	group.Use(middleware.JWTAuth())
	initAuthRouter(group)
	return router
}

func initNoAuthRouter(group *gin.RouterGroup) {
	group.POST("/register", register)
	group.POST("/login", login)
}

func initAuthRouter(group *gin.RouterGroup) {
	initAuthUserRouter(group)
	initAuthClassRouter(group)
	initAuthVideoRouter(group)
	initAuthCommentRouter(group)
}

func initAuthUserRouter(group *gin.RouterGroup) {
	// 显示单个用户信息
	// 删除单个用户信息(注销，此时用户下的分类、视频和评论都需要删除:定时任务),这里只需要打标记即可
	// 修改单个用户信息

	// 显示所有用户信息(管理员用)
	user := group.Group("/user")
	{
		user.GET("/:user_name", getUserInfo)
		user.DELETE("/:user_name", deleteUserInfo)
		user.PUT("/:user_name", updateUserInfo)

		user.GET("/", getAllUserInfo)
	}
}

func initAuthClassRouter(group *gin.RouterGroup) {

}

func initAuthVideoRouter(group *gin.RouterGroup) {

}

func initAuthCommentRouter(group *gin.RouterGroup) {

}
