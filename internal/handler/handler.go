package handler

import "github.com/kozhamseitova/api-blog/internal/service"

type Handler struct {
	srvs service.Service
}

func New(srvs service.Service) *Handler {
	return &Handler{
		srvs: srvs,
	}
}
