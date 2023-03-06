package auth

import (
	"context"

	"github.com/florentinuskev/simple-todo/internal/dto"
)

type AuthService interface {
	GetProfile(c context.Context, userReq *dto.GetProfileReq) (*dto.GetProfileRes, error)
	UserRegister(c context.Context, userReq *dto.UserRegisterReq) (*dto.UserRegisterRes, error)
	UserLogin(c context.Context, userReq *dto.UserLoginReq) (*dto.UserLoginRes, error)
}
