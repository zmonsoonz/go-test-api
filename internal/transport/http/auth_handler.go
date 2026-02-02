package transport_http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zmonsoonz/go-test-api/internal/content/domain"
	"github.com/zmonsoonz/go-test-api/internal/content/ports"
	"github.com/zmonsoonz/go-test-api/internal/platform/httpx"
)

type AuthHandler struct {
	authService ports.AuthService
}

func NewAuthHandler(authService ports.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) InitAuthRoutes(r gin.IRouter) {
	auth := r.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}
}
func (h *AuthHandler) SignUp (c *gin.Context) {
	var input domain.User

	if err := c.BindJSON(&input); err != nil {
		httpx.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
         
	id, err := h.authService.SignUp(input)
	if err != nil {
		httpx.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
func (h *AuthHandler) SignIn (c *gin.Context)  {
	var input SignInInput

	if err := c.BindJSON(&input); err != nil {
		httpx.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
         
	token, err := h.authService.SignIn(input.Username, input.Password)
	if err != nil {
		httpx.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}