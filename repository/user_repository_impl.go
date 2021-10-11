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

func (repository *UserRepositoryImpl) Save(ctx context.Context, user domain.User) (domain.User, error) {
	return user, repository.server.DB.Create(&user).Error
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, user domain.User) (domain.User, error) {
	if err := repository.server.DB.First(&user, user.ID).Error; err != nil {
		return user, err
	}
	return user, repository.server.DB.Model(&user).Updates(&user).Error
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, user domain.User, userID uint) error {
	if err := repository.server.DB.First(&user, user.ID).Error; err != nil {
		return err
	}
	return repository.server.DB.Model(&user).Updates(&user).Error
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, user domain.User, userID uint) (domain.User, error) {
	return user, repository.server.DB.First(&user, user.ID).Error
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, users []domain.User) ([]domain.User, error) {
	return users, repository.server.DB.First(&users).Error

}
