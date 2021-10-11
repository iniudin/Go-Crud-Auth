package service

import (
	"context"
	"go-crud-auth/model/web/requests"
	"go-crud-auth/model/web/responses"
)

type UserService interface {
	Create(ctx context.Context, request requests.UserCreate) responses.User
	Update(ctx context.Context, request requests.UserUpdate) responses.User
	Delete(ctx context.Context, userId uint)
	FindById(ctx context.Context, userId uint) responses.User
	FindAll(ctx context.Context) []responses.User
}
