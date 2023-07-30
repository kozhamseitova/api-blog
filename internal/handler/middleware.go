package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/kozhamseitova/api-blog/api"
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
		err := errors.New("authorization header is not set")
		c.AbortWithStatusJSON(http.StatusUnauthorized, &api.Error{
			Code:    -1,
			Message: err.Error(),
		})
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		err := errors.New("authorization header incorrect format")
		c.AbortWithStatusJSON(http.StatusUnauthorized, &api.Error{
			Code:    -2,
			Message: err.Error(),
		})
		return
	}

	userId, err := h.srvs.VerifyToken(headerParts[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, &api.Error{
			Code:    -3,
			Message: err.Error(),
		})
		return
	}

	c.Set(userCtx, userId)
}
