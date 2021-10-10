package repository

import (
	"context"
	"go-crud-auth/app"
	"go-crud-auth/model/domain"
)

type UserRepositoryImpl struct {
	server *app.Server
}

func NewUserRepository(server *app.Server) UserRepository {
	return &UserRepositoryImpl{server: server}
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
