package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"webapp/controller"
	"webapp/logger"
	"webapp/middleware"
)

func Setup()*gin.Engine{
	mode := viper.GetString("app.mode")
	if mode == "dev"{
		gin.SetMode(gin.DebugMode)
	}else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(),logger.GinRecovery(true))
	r.POST("/signup",controller.SignUpHandler)
	r.POST("/login",controller.LoginHandler)
	r.POST("/index",middleware.JWTAuthMiddleware(),controller.Index)
	return r
}


