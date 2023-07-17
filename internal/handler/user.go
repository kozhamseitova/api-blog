package handler

import (
	"api-blog/internal/entity"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (h *Handler) createUser(ctx *gin.Context) {
	var req entity.User

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	err = h.srvs.CreateUser(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *Handler) loginUser(ctx *gin.Context) {
	var req entity.UserLogin

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	token, err := h.srvs.Login(ctx, req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (h *Handler) updateUser(ctx *gin.Context) {
	var req entity.User

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

	req.ID = userId.(int64)

	err = h.srvs.UpdateUser(ctx, &req)
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

func (h *Handler) deleteUser(ctx *gin.Context) {
	id, ok := ctx.Get(userCtx)
	if !ok {
		ctx.JSON(http.StatusForbidden, &Error{
			Code:    http.StatusForbidden,
			Message: "user not found",
		})
		return
	}

	err := h.srvs.DeleteUser(ctx, id.(int64))
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
