package service

import (
	"go-crud-auth/config"
	"go-crud-auth/helper"
	"go-crud-auth/model/domain"
	"time"

	"github.com/golang-jwt/jwt"
)

type AuthServiceImpl struct {
	config *config.Config
}

func NewAuthService(config *config.Config) AuthService {
	return &AuthServiceImpl{
		config: config,
	}
}

func (service *AuthServiceImpl) CreateAccessToken(user domain.User) (accessToken string, expires int64, err error) {
	exp := time.Now().Add(time.Hour * ExpiresCount).Unix()
	claims := &JWTCustomClaims{
		Username:       user.Username,
		Id:             user.ID,
		StandardClaims: jwt.StandardClaims{ExpiresAt: exp},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(service.config.JWTSecretKey))
	helper.PanicError(err)

	return t, expires, err
}
