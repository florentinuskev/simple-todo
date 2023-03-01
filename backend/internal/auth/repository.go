package auth

import (
	"context"

	"github.com/florentinuskev/simple-todo/internal/dao"
)

type AuthRepository interface {
	CreateUser(c context.Context, u *dao.User) (*dao.User, error)
	FindUserById(c context.Context, uid string) (*dao.User, error)
	FindUserByUsername(c context.Context, username string) (*dao.User, error)
}
