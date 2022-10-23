package routers

import (
	"github.com/gin-gonic/gin"
	"video_server/global"

	ginprometheus "github.com/zinclabs/go-gin-prometheus"
)

// SetPrometheus sets up prometheus metrics for giin
func SetPrometheus(app *gin.Engine) {
	if !global.ServerConfig.PrometheusEnable {
		return
	}

	p := ginprometheus.NewPrometheus("Twitta", []*ginprometheus.Metric{})
	p.Use(app)
}
