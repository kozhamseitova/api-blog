package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kozhamseitova/api-blog/api"
	"github.com/kozhamseitova/api-blog/internal/entity"
	"log"
	"net/http"
)

func (h *Handler) createUser(ctx *gin.Context) {
	var req entity.User

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	err = h.srvs.CreateUser(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &api.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *Handler) loginUser(ctx *gin.Context) {
	var req api.LoginRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &api.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	token, err := h.srvs.Login(ctx, req.Username, req.Password)
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
		Data:    token,
	})
}

func (h *Handler) updateUser(ctx *gin.Context) {
	var req entity.User

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

	req.ID = userId

	err = h.srvs.UpdateUser(ctx, &req)
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

func (h *Handler) deleteUser(ctx *gin.Context) {
	id, ok := ctx.MustGet(userCtx).(int64)
	if !ok {
		log.Printf("can't get userID")
		ctx.JSON(http.StatusForbidden, &api.Error{
			Code:    http.StatusForbidden,
			Message: "can't get userID from auth",
		})
		return
	}

	err := h.srvs.DeleteUser(ctx, id)
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
		Data:    id,
	})
}
