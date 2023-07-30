package service

import (
	"github.com/kozhamseitova/api-blog/internal/config"
	"github.com/kozhamseitova/api-blog/internal/repository"
	"github.com/kozhamseitova/api-blog/pkg/jwttoken"
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
