package service

import (
	"api-blog/internal/config"
	"api-blog/internal/repository"
	"api-blog/pkg/jwttoken"
)

type Manager struct {
	Repository repository.Repository
	Config     *config.Config
	Token      *jwttoken.JWTToken
}

func New(repository repository.Repository, config *config.Config, token *jwttoken.JWTToken) *Manager {
	return &Manager{
		Repository: repository,
		Config:     config,
		Token:      token,
	}
}
