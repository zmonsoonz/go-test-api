package bootstrap

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func InitRouter(handlers *Handlers) *gin.Engine {
	router := gin.New()

	handlers.Auth.InitAuthRoutes(router)

	router.Group("/api", handlers.Auth.UserIdentity)
	return router
}