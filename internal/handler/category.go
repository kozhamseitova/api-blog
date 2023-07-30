package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kozhamseitova/api-blog/api"
	"net/http"
)

func (h *Handler) getAllCategories(ctx *gin.Context) {

	categories, err := h.srvs.GetCategories(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.Ok{
		Code:    0,
		Message: "successfully updated",
		Data:    categories,
	})
}
