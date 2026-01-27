package bootstrap

import (
	"github.com/gin-gonic/gin"
)


func InitRouter(handlers *Handlers) *gin.Engine {
	router := gin.New()
	{
		handlers.Auth.InitAuthRoutes(router)
	}
	return router
}