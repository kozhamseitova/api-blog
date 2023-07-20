package jwttoken

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

var (
	ErrInvalidToken = errors.New("invalid token err")
	ErrExpiredToken = errors.New("expired token err")
)

type JWTToken struct {
	secretKey string
}

func New(secretKey string) *JWTToken {
	return &JWTToken{
		secretKey: secretKey,
	}
}

func (j *JWTToken) CreateToken(userId int64, duration time.Duration) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("uuid new random err: %w", err)
	}

	payload := &JWTPayload{
		ID:        id,
		UserId:    userId,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return jwtToken.SignedString([]byte(j.secretKey))
}

func (j *JWTToken) ValidateToken(token string) (*JWTPayload, error) {
	keyFunc := func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if ok {
			return []byte(j.secretKey), nil
		}

		return nil, ErrInvalidToken
	}

	jwtToken, err := jwt.ParseWithClaims(token, &JWTPayload{}, keyFunc)
	if err != nil {
		validationErr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(validationErr, ErrExpiredToken) {
			return nil, ErrInvalidToken
		}
	}

	payload, ok := jwtToken.Claims.(*JWTPayload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
