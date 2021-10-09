package repository

import (
	"Secreto/model/domain"
	"context"
)

type UserRepository interface {
	Save(ctx context.Context, user domain.User)
	Update(ctx context.Context, user domain.User) (domain.User, error)
	Delete(ctx context.Context, uid uint)
	FindById(ctx context.Context, uid uint) (domain.User, error)
	FindAll(ctx context.Context) ([]domain.User, error)
}
