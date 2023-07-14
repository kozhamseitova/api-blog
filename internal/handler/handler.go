package handler

import "api-blog/internal/service"

type Handler struct {
	srvc service.Service
}

func New(srvc service.Service) *Handler {
	return &Handler{
		srvc: srvc,
	}
}
