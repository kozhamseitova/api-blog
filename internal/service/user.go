package service

import (
	"api-blog/internal/entity"
	"api-blog/pkg/util"
	"context"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
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
	user, err := m.Repository.Login(ctx, username, password)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	token, err := generateToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (m *Manager) UpdateUser(ctx context.Context, u *entity.User) error {
	err := m.Repository.UpdateUser(ctx, u)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) DeleteUser(ctx context.Context, id int64) error {
	err := m.Repository.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) VerifyToken(token string) error {
	return nil
}

func generateToken(id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		id,
	})

	return token.SignedString([]byte(signingKey))
}
