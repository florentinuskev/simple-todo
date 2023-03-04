package service

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"testing"
	"time"

	mock_auth "github.com/florentinuskev/simple-todo/internal/auth/mock"
	"github.com/florentinuskev/simple-todo/internal/dao"
	"github.com/florentinuskev/simple-todo/internal/dto"
	"github.com/florentinuskev/simple-todo/public/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

type eqCreateUserParamsMatcher struct {
	user     *dao.User
	password string
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(*dao.User)
	if !ok {
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(arg.Password), []byte(e.password))
	if err != nil {
		return false
	}
	e.user.ID = arg.ID
	e.user.Password = arg.Password

	return reflect.DeepEqual(e.user, arg)
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.user, e.password)
}

func EqCreateUserParams(user *dao.User, password string) gomock.Matcher {
	return eqCreateUserParamsMatcher{user, password}
}

func TestUserLoginService(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockAuthRepo := mock_auth.NewMockAuthRepository(ctrl)
	authSvc := NewAuthService(&utils.Config{}, mockAuthRepo)

	ctx := context.Background()

	user := &dao.User{
		Username: "test123",
		Password: "test123",
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	require.NoError(t, err)

	mockUser := &dao.User{
		Username: "test123",
		Password: string(hashedPassword),
		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	mockAuthRepo.EXPECT().FindUserByUsername(ctx, gomock.Eq(user.Username)).Return(mockUser, nil)

	userRes, err := authSvc.UserLogin(ctx, &dto.UserLoginReq{Username: user.Username, Password: user.Password})

	require.NoError(t, err)
	require.NotNil(t, userRes)
	require.NotEmpty(t, userRes.Token)
	require.Nil(t, err)
}

func TestUserLoginInvalidPass(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockAuthRepo := mock_auth.NewMockAuthRepository(ctrl)
	authSvc := NewAuthService(&utils.Config{}, mockAuthRepo)

	ctx := context.Background()

	user := &dao.User{
		Username: "test123",
		Password: "test123",
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	require.NoError(t, err)

	mockUser := &dao.User{
		Username: "test123",
		Password: string(hashedPassword),
		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	mockAuthRepo.EXPECT().FindUserByUsername(ctx, gomock.Eq(user.Username)).Return(mockUser, nil)

	userRes, err := authSvc.UserLogin(ctx, &dto.UserLoginReq{Username: user.Username, Password: "test12"})

	require.Error(t, err)
	require.Equal(t, uint32(400), userRes.Status)
	require.NotNil(t, userRes)
}

func TestUserRegisterService(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockAuthRepo := mock_auth.NewMockAuthRepository(ctrl)
	authSvc := NewAuthService(&utils.Config{}, mockAuthRepo)

	ctx := context.Background()

	user := &dao.User{
		Username: "test123",
		Password: "test123",
	}

	mockAuthRepo.EXPECT().FindUserByUsername(ctx, gomock.Eq(user.Username)).Return(nil, nil)
	mockAuthRepo.EXPECT().CreateUser(ctx, EqCreateUserParams(user, user.Password)).Return(user, nil)

	userRes, err := authSvc.UserRegister(ctx, &dto.UserRegisterReq{
		Username: user.Username,
		Password: user.Password,
	})

	t.Log(userRes.User)

	require.NoError(t, err)
	require.NotNil(t, userRes)
	require.Equal(t, user.Username, userRes.User.Username)
	require.Equal(t, uint32(201), userRes.Status)
}
