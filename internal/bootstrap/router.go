package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/zmonsoonz/go-test-api/internal/transport/middleware"
)


func InitRouter(handlers *Handlers, services *Services) *gin.Engine {
	router := gin.New()
	handlers.Auth.InitAuthRoutes(router)

	api := router.Group("/api")
	api.Use(middleware.UserIdentity(services.Auth))
	{

		handlers.Track.InitTrackRoutes(api)
	}
	
	return router
}