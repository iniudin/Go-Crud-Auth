package service

import (
	"context"
	"go-crud-auth/helper"
	"go-crud-auth/model/domain"
	"go-crud-auth/model/web/requests"
	"go-crud-auth/model/web/responses"
	"go-crud-auth/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, request requests.UserCreate) responses.User {
	userRequest := domain.User{
		Name:     request.Name,
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}
	user, err := service.UserRepository.Save(ctx, userRequest)
	helper.PanicError(err)

	return responses.User{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
	}
}

func (service *UserServiceImpl) Update(ctx context.Context, request requests.UserUpdate) responses.User {
	panic("not implemented") // TODO: Implement
}

func (service *UserServiceImpl) Delete(ctx context.Context, userId uint) {
	panic("not implemented") // TODO: Implement
}

func (service *UserServiceImpl) FindById(ctx context.Context, userId uint) responses.User {
	panic("not implemented") // TODO: Implement
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []responses.User {
	panic("not implemented") // TODO: Implement
}
