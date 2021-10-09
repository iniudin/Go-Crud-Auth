package repository

import (
	"Secreto/app"
	"Secreto/model/domain"
	"context"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	connection *gorm.DB
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{connection: app.NewDatabase()}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, user domain.User) {
	panic("not implemented") // TODO: Implement
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, user domain.User) (domain.User, error) {
	panic("not implemented") // TODO: Implement
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, uid uint) {
	panic("not implemented") // TODO: Implement
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, uid uint) (domain.User, error) {
	panic("not implemented") // TODO: Implement
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context) ([]domain.User, error) {
	panic("not implemented") // TODO: Implement
}
