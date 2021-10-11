package repository

import (
	"context"
	"go-crud-auth/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, user domain.User) (domain.User, error)
	Update(ctx context.Context, user domain.User) (domain.User, error)
	Delete(ctx context.Context, user domain.User, userID uint) error
	FindById(ctx context.Context, user domain.User, userID uint) (domain.User, error)
	FindAll(ctx context.Context, users []domain.User) ([]domain.User, error)
}
