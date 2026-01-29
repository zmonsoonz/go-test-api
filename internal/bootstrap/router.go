package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/zmonsoonz/go-test-api/internal/transport/middleware"
)


func InitRouter(handlers *Handlers, services *Services) *gin.Engine {
	router := gin.New()
	router.Use(middleware.UserIdentity(services.Auth))

	router.Group("/api")
	handlers.Auth.InitAuthRoutes(router)

	return router
}