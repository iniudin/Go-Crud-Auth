package repository

import "context"

type AuthRepository interface {
	Login(ctx context.Context, username, password string)
	Logout(ctx context.Context, loginSession string)
}
