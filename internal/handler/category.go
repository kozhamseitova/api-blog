package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getAllCategories(ctx *gin.Context) {

	categories, err := h.srvs.GetCategories(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": categories,
	})
}
