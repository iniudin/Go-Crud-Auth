package service

import (
	"Secreto/helper"
	"Secreto/model/domain"
	"time"

	"github.com/golang-jwt/jwt"
)

type AuthServiceImpl struct {
}

func NewAuthService() AuthService {
	return &AuthServiceImpl{}
}

func (service *AuthServiceImpl) CreateAccessToken(user domain.User) (accessToken string, expires int64, err error) {
	exp := time.Now().Add(time.Hour * ExpiresCount).Unix()
	claims := &JWTCustomClaims{
		Username:       user.Username,
		Id:             user.ID,
		StandardClaims: jwt.StandardClaims{ExpiresAt: exp},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secreto"))
	helper.PanicError(err)

	return t, expires, err
}
