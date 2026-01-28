package transport_http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zmonsoonz/go-test-api/internal/content/domain"
	"github.com/zmonsoonz/go-test-api/internal/content/ports"
	"github.com/zmonsoonz/go-test-api/internal/platform/httpx"
)

type Handler struct {
	authService ports.AuthService
}

func NewAuthHandler(authService ports.AuthService) *Handler {
	return &Handler{authService: authService}
}

func (h *Handler) InitAuthRoutes(r gin.IRouter) {
	auth := r.Group("auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}
}
func (h *Handler) SignUp (c *gin.Context) {
	var input domain.User

	if err := c.BindJSON(&input); err != nil {
		httpx.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
         
	id, err := h.authService.CreateUser(input)
	if err != nil {
		httpx.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
func (h *Handler) SignIn (c *gin.Context)  {
	var input SignInInput

	if err := c.BindJSON(&input); err != nil {
		httpx.NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
         
	token, err := h.authService.GenerateToken(input.Username, input.Password)
	if err != nil {
		httpx.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}