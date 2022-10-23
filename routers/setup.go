package routers

import (
	"github.com/gin-gonic/gin"
	"video_server/global"
)

func Setup(app *gin.Engine) {
	gin.SetMode(global.ServerConfig.DebugMode)
	app.Use(gin.Logger(), gin.Recovery())

	// Debug for gin
	if gin.Mode() == gin.DebugMode {
		SetPProf(app)
	}
	SetPrometheus(app) // Set up Prometheus.
	SetRouters(app)    // Set up all API routers.
}
