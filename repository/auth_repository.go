package repository

import "context"

type AuthRepository interface {
	Login(ctx context.Context, username, password string)
}
