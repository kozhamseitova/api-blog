package jwttoken

import (
	"github.com/google/uuid"
	"time"
)

type JWTPayload struct {
	ID        uuid.UUID `json:"id"`
	UserId    int64     `json:"user_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func (j JWTPayload) Valid() error {
	if time.Now().After(j.ExpiredAt) {
		return ErrExpiredToken
	}

	return nil
}
