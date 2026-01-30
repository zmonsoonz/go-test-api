package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zmonsoonz/go-test-api/internal/content/ports"
	"github.com/zmonsoonz/go-test-api/internal/platform/httpx"
)

const (
	AuthHeader = "Authorization"
)

func UserIdentity(authService ports.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader(AuthHeader)
		if header == "" {
			httpx.NewErrorResponse(c, http.StatusUnauthorized, "authorization header is empty")
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			httpx.NewErrorResponse(c, http.StatusUnauthorized, "invalid authorization header format")
			return
		}

		userID, err := authService.ParseToken(headerParts[1])
		if err != nil {
			httpx.NewErrorResponse(c, http.StatusUnauthorized, "invalid token")
			return
		}

		c.Set("userId", userID)
		c.Next()
	}
}

func GetUserId(c *gin.Context) (int, error) {
	userId, ok := c.Get("userId")
	if !ok {
		httpx.NewErrorResponse(c, http.StatusUnauthorized, "user not authorized")
		return 0, errors.New("user not authorized")
	}
		
	id, ok := userId.(int) // приводим к типу int т.к. Get возвращает interface{}
	if !ok {
		httpx.NewErrorResponse(c, http.StatusUnauthorized, "user not authorized")
		return 0, errors.New("user not authorized")
	}

	return id, nil
}