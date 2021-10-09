package repository

import "context"

type AuthRepositoryImpl struct {
}

func NewAuthRepository() AuthRepository {
	return &AuthRepositoryImpl{}
}

func (repository *AuthRepositoryImpl) Login(ctx context.Context, username string, password string) {
	panic("not implemented") // TODO: Implement
}

func (repository *AuthRepositoryImpl) Logout(ctx context.Context, loginSession string) {
	panic("not implemented") // TODO: Implement
}
