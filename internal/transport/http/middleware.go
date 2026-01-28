package transport_http

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	AuthHeader = "Authorization"
)

func (h *Handler) UserIdentity(c *gin.Context) {
	header := c.GetHeader(AuthHeader)
	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "no token")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid token format")
		return
	}

	userId, err := h.authService.ParseToken(headerParts[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid token")
		return
	}

	c.Set("userId", userId)
}