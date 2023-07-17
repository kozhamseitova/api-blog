package handler

import (
	"api-blog/internal/entity"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) createArticle(ctx *gin.Context) {
	var req entity.Article

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	userId, ok := ctx.Get(userCtx)
	if !ok {
		ctx.JSON(http.StatusForbidden, &Error{
			Code:    http.StatusForbidden,
			Message: "user not found",
		})
		return
	}

	req.UserID = userId.(int64)

	err = h.srvs.CreateArticle(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *Handler) updateArticle(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &Error{
			Code:    http.StatusBadRequest,
			Message: "invalid id param",
		})
		return
	}

	var req entity.Article

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	req.ID = int64(id)

	userId, ok := ctx.Get(userCtx)
	if !ok {
		ctx.JSON(http.StatusForbidden, &Error{
			Code:    http.StatusForbidden,
			Message: "user not found",
		})
		return
	}

	req.UserID = userId.(int64)

	err = h.srvs.UpdateArticle(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "successfully updated",
	})
}

func (h *Handler) deleteArticle(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &Error{
			Code:    http.StatusBadRequest,
			Message: "invalid id param",
		})
		return
	}

	userId, ok := ctx.Get(userCtx)
	if !ok {
		ctx.JSON(http.StatusForbidden, &Error{
			Code:    http.StatusForbidden,
			Message: "user not found",
		})
		return
	}

	err = h.srvs.DeleteArticle(ctx, int64(id), userId.(int64))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "successfully deleted",
	})
}

func (h *Handler) getArticleById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &Error{
			Code:    http.StatusBadRequest,
			Message: "invalid id param",
		})
		return
	}

	article, err := h.srvs.GetArticleByID(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": article,
	})
}

func (h *Handler) getAllArticles(ctx *gin.Context) {

	articles, err := h.srvs.GetAllArticles(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": articles,
	})
}

func (h *Handler) getArticleByUserId(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &Error{
			Code:    http.StatusBadRequest,
			Message: "invalid id param",
		})
		return
	}

	articles, err := h.srvs.GetArticlesByUserID(ctx, int64(userId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": articles,
	})
}
