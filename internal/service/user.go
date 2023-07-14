package service

import (
	"api-blog/internal/entity"
	"api-blog/pkg/util"
	"context"
)

func (m *Manager) CreateUser(ctx context.Context, u *entity.User) error {
	hashedPassword, err := util.HashPassword(u.Password)
}
