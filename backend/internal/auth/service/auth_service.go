package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/florentinuskev/simple-todo/internal/auth"
	"github.com/florentinuskev/simple-todo/internal/dao"
	"github.com/florentinuskev/simple-todo/internal/dto"
	"github.com/florentinuskev/simple-todo/public/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	cfg *utils.Config
	r   auth.AuthRepository
}

func NewAuthService(cfg *utils.Config, r auth.AuthRepository) auth.AuthService {
	return &AuthService{cfg: cfg, r: r}
}

func (as *AuthService) UserRegister(c context.Context, userReq *dto.UserRegisterReq) (*dto.UserRegisterRes, error) {
	existUser, err := as.r.FindUserByUsername(c, userReq.Username)

	if existUser != nil {
		return &dto.UserRegisterRes{
			Status: 400,
		}, errors.New("username exists")
	}

	if err != sql.ErrNoRows && err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), 10)

	if err != nil {
		return nil, err
	}
	createdUser, err := as.r.CreateUser(c, &dao.User{Username: userReq.Username, Password: string(hashedPassword)})

	if err != nil {
		return nil, err
	}

	return &dto.UserRegisterRes{
		Status: 201,
		User:   createdUser,
	}, nil
}

func (as *AuthService) UserLogin(c context.Context, userReq *dto.UserLoginReq) (*dto.UserLoginRes, error) {
	user, err := as.r.FindUserByUsername(c, userReq.Username)

	if user == nil {
		return nil, errors.New("Username does not exists")
	}

	if err != sql.ErrNoRows && err != nil {
		return nil, err
	}

	errCompare := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userReq.Password))

	if errCompare != nil {
		if errCompare == bcrypt.ErrMismatchedHashAndPassword {
			return &dto.UserLoginRes{
				Status: 400,
			}, errCompare
		}
		return nil, errCompare
	}

	tokenString, err := utils.GenerateJWT(user, as.cfg.Env["JWT_SECRET"], 20)

	if err != nil {
		return nil, err
	}

	return &dto.UserLoginRes{
		Status: 200,
		Token:  tokenString,
	}, nil
}
func (as *AuthService) GetProfile(c context.Context, userReq *dto.GetProfileReq) (*dto.GetProfileRes, error) {

	user, err := as.r.FindUserById(c, userReq.UID)

	if err != nil {
		return nil, err
	}

	return &dto.GetProfileRes{
		Status: 200,
		User: &dao.User{
			ID:        user.ID,
			Username:  user.Username,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}, nil
}
