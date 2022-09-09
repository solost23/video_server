package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "video_server/docs" // 必须要导入生成的docs文档包
	"video_server/pkg/middleware"
)

type ErrCode int

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	group := router.Group("api")
	initNoAuthRouter(group)
	// 注意 role 需要再思考一下，不一定要放在这里
	//group.Use(jwt.JWTAuth, role.CheckRole)
	group.Use(middleware.JWTAuth())
	initAuthRouter(group)
	return router
}

func initNoAuthRouter(group *gin.RouterGroup) {
	group.POST("register", register)
	group.POST("register/avatar", uploadAvatar)
	group.POST("login", login)
	group.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func initAuthRouter(group *gin.RouterGroup) {
	initAuthUserRouter(group)
	initAuthCategoryRouter(group)
	initAuthVideoRouter(group)
	initAuthCommentRouter(group)
	initAuthRoleRouter(group)
}

func initAuthUserRouter(group *gin.RouterGroup) {
	user := group.Group("user")
	{
		// 注销用户
		user.POST("logout", logout)
		// 显示单个用户信息
		user.GET(":id", userDetail)
		// 删除单个用户信息(注销，此时用户下的分类、视频和评论都需要删除:定时任务),这里只需要打标记即可
		user.DELETE(":id", userDelete)
		// 修改单个用户信息
		user.PUT(":id", userUpdate)
		// 显示所有用户信息
		user.GET("", userList)
	}
}

func initAuthCategoryRouter(group *gin.RouterGroup) {
	class := group.Group("category")
	{
		class.POST("", categoryInsert)
		class.PUT(":id", categoryUpdate)
		class.GET("", categoryList)
	}
}

func initAuthVideoRouter(group *gin.RouterGroup) {
	video := group.Group("video")
	{
		// 上传视频图片接口
		video.POST("img", videoUploadImg)
		// 上传视频流接口
		video.POST("vid", videoUploadVid)
		// 提交视频信息接口
		video.POST("", videoInsert)
		// 删除就是将video信息的delete_status的字段修改为已删除
		video.DELETE(":id", videoDelete)
		// 获取单个视频信息(视频流直接就可以通过video_url字段访问到，所以不用处理文件)
		video.GET(":id", videoDetail)
		// 首页 支持获取所有视频，支持按照 分类名，用户名，视频标题 搜索，并支持分页操作
		video.GET("", videoList)
	}
}

func initAuthCommentRouter(group *gin.RouterGroup) {
	comment := group.Group("comment")
	{
		comment.POST("", commentCreate)
		comment.DELETE(":id", commentDelete)
		comment.GET("", commentList)
	}
}

func initAuthRoleRouter(group *gin.RouterGroup) {
	role := group.Group("role")
	{
		role.POST("", roleInsert)
		role.DELETE(":id", roleDelete)
		role.GET("", roleList)
	}
}
