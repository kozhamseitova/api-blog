package service

import (
	"api-blog/internal/entity"
	"api-blog/pkg/util"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

const signingKey = "kwjebr23oif99we"

type tokenClaims struct {
	jwt.StandardClaims
	UserId int64 `json:"user_id"`
}

func (m *Manager) CreateUser(ctx context.Context, u *entity.User) error {
	hashedPassword, err := util.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword

	err = m.Repository.CreateUser(ctx, u)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) Login(ctx context.Context, username, password string) (string, error) {
	user, err := m.Repository.GetUserByUsernameAndPassword(ctx, username, password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("user not found")
		}
		return "", fmt.Errorf("get user err: %w", err)
	}

	err = util.CheckPassword(password, user.Password)
	if err != nil {
		return "", fmt.Errorf("incorrect password: %w", err)
	}

	token, err := m.Token.CreateToken(user.ID, m.Config.Token.TimeToLive)

	if err != nil {
		return "", fmt.Errorf("create token err: %w", err)
	}

	return token, nil
}

func (m *Manager) UpdateUser(ctx context.Context, u *entity.User) error {
	return m.Repository.UpdateUser(ctx, u)
}

func (m *Manager) DeleteUser(ctx context.Context, id int64) error {
	return m.Repository.DeleteUser(ctx, id)
}

func (m *Manager) VerifyToken(token string) (int64, error) {
	payload, err := m.Token.ValidateToken(token)
	if err != nil {
		return 0, fmt.Errorf("validate token err: %w", err)
	}

	return payload.UserId, nil
}
