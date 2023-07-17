package handler

import (
	"fmt"
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
			Code:    -1,
			Message: "empty auth header",
		})
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		c.JSON(http.StatusUnauthorized, &Error{
			Code:    -1,
			Message: "invalid auth header",
		})
		return
	}

	userId, err := h.srvs.ParseToken(headerParts[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, &Error{
			Code:    -1,
			Message: "token parse error",
		})
		return
	}

	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		c.JSON(http.StatusInternalServerError, &Error{
			Code:    -1,
			Message: "user id not found",
		})
		return 0, fmt.Errorf("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, &Error{
			Code:    -1,
			Message: "invalid type user id",
		})
		return 0, fmt.Errorf("invalid type user id")
	}

	return idInt, nil
}
