package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == " " {
		c.JSON(http.StatusUnauthorized, &Error{
			Code:    http.StatusUnauthorized,
			Message: "empty auth header",
		})
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		c.JSON(http.StatusUnauthorized, &Error{
			Code:    http.StatusUnauthorized,
			Message: "invalid auth header",
		})
		return
	}

	userId, err := h.srvs.ParseToken(headerParts[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, &Error{
			Code:    http.StatusUnauthorized,
			Message: "token parse error",
		})
		return
	}

	c.Set(userCtx, userId)
}
