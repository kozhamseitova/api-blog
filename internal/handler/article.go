package handler

import (
	"api-blog/api"
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
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	userId, ok := ctx.MustGet(userCtx).(int64)
	if !ok {
		log.Printf("can't get userID")
		ctx.JSON(http.StatusForbidden, &api.Error{
			Code:    http.StatusForbidden,
			Message: "can't get userID from auth",
		})
		return
	}

	req.UserID = userId

	err = h.srvs.CreateArticle(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *Handler) updateArticle(ctx *gin.Context) {
	var req entity.Article

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	req.ID = int64(req.ID)

	userId, ok := ctx.MustGet(userCtx).(int64)
	if !ok {
		log.Printf("can't get userID")
		ctx.JSON(http.StatusForbidden, &api.Error{
			Code:    http.StatusForbidden,
			Message: "can't get userID from auth",
		})
		return
	}

	req.UserID = userId

	err = h.srvs.UpdateArticle(ctx, &req)
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
		Data:    userId,
	})
}

func (h *Handler) deleteArticle(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    http.StatusBadRequest,
			Message: "invalid id param",
		})
		return
	}

	userId, ok := ctx.MustGet(userCtx).(int64)
	if !ok {
		log.Printf("can't get userID")
		ctx.JSON(http.StatusForbidden, &api.Error{
			Code:    http.StatusForbidden,
			Message: "can't get userID from auth",
		})
		return
	}

	err = h.srvs.DeleteArticle(ctx, int64(id), userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.Ok{
		Code:    0,
		Message: "successfully deleted",
		Data:    userId,
	})
}

func (h *Handler) getArticleById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    http.StatusBadRequest,
			Message: "invalid id param",
		})
		return
	}

	article, err := h.srvs.GetArticleByID(ctx, int64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.Ok{
		Code:    0,
		Message: "success",
		Data:    article,
	})
}

func (h *Handler) getAllArticles(ctx *gin.Context) {

	articles, err := h.srvs.GetAllArticles(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.Error{
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
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    http.StatusBadRequest,
			Message: "invalid id param",
		})
		return
	}

	articles, err := h.srvs.GetArticlesByUserID(ctx, int64(userId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &api.Ok{
		Code:    0,
		Message: "success",
		Data:    articles,
	})
}
