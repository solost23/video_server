package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "video_server/docs" // 必须要导入生成的docs文档包
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
	group.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func initAuthRouter(group *gin.RouterGroup) {
	initAuthUserRouter(group)
	initAuthClassRouter(group)
	initAuthVideoRouter(group)
	initAuthCommentRouter(group)
	initAuthRoleRouter(group)
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

		user.GET("", getAllUserInfo)
	}
}

func initAuthClassRouter(group *gin.RouterGroup) {
	class := group.Group("/class")
	{
		class.POST("/:user_name", createClass)
		class.PUT("/:user_name/:class_id", updateClass)
		class.GET("/:user_name", getUserAllClass)
	}
}

func initAuthVideoRouter(group *gin.RouterGroup) {
	video := group.Group("/video")
	{
		// 提交视频信息,通过表单，视频流和视频信息一起上传
		video.POST("/:user_name/:class_id", createVideo)
		// 删除就是将video信息的delete_status的字段修改为已删除
		video.DELETE("/:user_name/:class_id/:video_id", deleteVideo)
		// 获取单个视频信息(视频流直接就可以通过video_url字段访问到，所以不用处理文件)
		video.GET("/:user_name/:class_id/:video_id", getVideo)
		// 获取此用户分类下所有视频
		video.GET("/:user_name/:class_id", getVideoByUserNameAndClassID)
		// 获取此用户下所有视频
		video.GET("/:user_name", getVideoByUserName)
		// 获取所有视频
		video.GET("", getALLVideo)
	}
}

func initAuthCommentRouter(group *gin.RouterGroup) {
	comment := group.Group("/comment")
	{
		comment.POST("/:video_id", createComment)
		comment.DELETE("/:video_id/:comment_id", deleteComment)
		comment.GET("/:video_id", getCommentByVideoID)
	}
}

func initAuthRoleRouter(group *gin.RouterGroup) {
	role := group.Group("/role")
	{
		role.POST("", addRoleAuth)
		role.DELETE("", deleteRoleAuth)
		role.GET("", getAllRoleAuth)
		role.GET("/:role_name", getRoleAuth)
	}
}
