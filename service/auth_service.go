package service

import (
	"go-crud-auth/model/domain"

	"github.com/golang-jwt/jwt"
)

const (
	ExpiresCount = 3
)

type JWTCustomClaims struct {
	Username string `json:"name"`
	Id       uint   `json:"id"`
	jwt.StandardClaims
}

type AuthService interface {
	CreateAccessToken(user domain.User) (accessToken string, expires int64, err error)
}
