package repository

import (
	"context"

	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{DB: db}
}

func (repository *AuthRepositoryImpl) Login(ctx context.Context, username string, password string) {
	panic("not implemented") // TODO: Implement
}
