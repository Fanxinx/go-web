package routes

import (
	"github.com/gin-gonic/gin"
	"webapp/logger"
)

func Setup()*gin.Engine{
	r := gin.New()
	r.Use(logger.GinLogger(),logger.GinRecovery(true))
	return r
}


