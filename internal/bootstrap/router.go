package bootsrap

import (
	"github.com/gin-gonic/gin"
	"github.com/zmonsoonz/go-test-api/internal/content/ports"
)

type Router struct {
}

type Entities struct {
	User ports.UserService
}

func (h *Router) InitRoutes(e Entities) *gin.Engine {
	router := gin.New()

	router.Group("/api/v1/")
	return router
}