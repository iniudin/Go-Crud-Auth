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

func (service *UserServiceImpl) Update(ctx context.Context, request requests.UserUpdate, userId uint) responses.User {
	userRequest := domain.User{
		Name:     request.Name,
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}
	userRequest.ID = userId

	user, err := service.UserRepository.Update(ctx, userRequest)
	helper.PanicError(err)

	return responses.User{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
	}
}

func (service *UserServiceImpl) Delete(ctx context.Context, userId uint) {
	user := domain.User{}
	user.ID = userId

	err := service.UserRepository.Delete(ctx, user)
	helper.PanicError(err)
}

func (service *UserServiceImpl) FindById(ctx context.Context, userId uint) responses.User {
	user := domain.User{}
	user.ID = userId

	user, err := service.UserRepository.FindById(ctx, user)
	helper.PanicError(err)

	return responses.User{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
	}
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []responses.User {
	var userResponses []responses.User
	users, err := service.UserRepository.FindAll(ctx, []domain.User{})
	helper.PanicError(err)
	for _, user := range users {
		userResponses = append(userResponses, responses.User{
			ID:       user.ID,
			Name:     user.Name,
			Username: user.Username,
			Email:    user.Email,
		})
	}

	return userResponses
}
