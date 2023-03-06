package controller

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mock_auth "github.com/florentinuskev/simple-todo/internal/auth/mock"
	"github.com/florentinuskev/simple-todo/internal/dao"
	"github.com/florentinuskev/simple-todo/internal/dto"
	"github.com/florentinuskev/simple-todo/public/utils"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestUserRegister(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthSvc := mock_auth.NewMockAuthService(ctrl)
	authController := NewAuthController(&utils.Config{}, mockAuthSvc)

	newUser := &dao.User{
		Username:  "test123",
		Password:  "test123",
		CreatedAt: sql.NullTime{},
		UpdatedAt: sql.NullTime{},
	}

	userReq := &dto.UserRegisterReq{Username: newUser.Username, Password: newUser.Password}

	buf, err := utils.AnyToBytesBuffer(userReq)

	require.NoError(t, err)
	require.NotNil(t, buf)
	require.Nil(t, err)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/register", buf)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	mockAuthSvc.EXPECT().UserRegister(gomock.Any(), gomock.Eq(userReq)).Return(&dto.UserRegisterRes{
		Status: 201,
		User:   newUser,
	}, nil)

	errRes := authController.UserRegister(c)

	require.NoError(t, errRes)
}

func TestUserLogin(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthSvc := mock_auth.NewMockAuthService(ctrl)
	authController := NewAuthController(&utils.Config{}, mockAuthSvc)

	user := &dao.User{
		ID:        uuid.NewString(),
		Username:  "test123",
		Password:  "test123",
		CreatedAt: sql.NullTime{},
		UpdatedAt: sql.NullTime{},
	}

	userReq := &dto.UserLoginReq{Username: user.Username, Password: user.Password}

	buf, err := utils.AnyToBytesBuffer(userReq)

	require.NoError(t, err)
	require.NotNil(t, buf)
	require.Nil(t, err)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", buf)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	tokenString, err := utils.GenerateJWT(user, "HELLOworld123", 20)

	require.NoError(t, err)
	require.NotEmpty(t, tokenString)

	mockAuthSvc.EXPECT().UserLogin(gomock.Any(), gomock.Eq(userReq)).Return(&dto.UserLoginRes{
		Status: 200,
		Token:  tokenString,
	}, nil)

	errRes := authController.UserLogin(c)

	require.NoError(t, errRes)
}

func TestGetProfile(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthSvc := mock_auth.NewMockAuthService(ctrl)
	authController := NewAuthController(&utils.Config{}, mockAuthSvc)

	user := &dao.User{
		ID:       uuid.NewString(),
		Username: "test123",
		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	userReq := &dto.GetProfileReq{UID: user.ID}

	buf, err := utils.AnyToBytesBuffer(userReq)

	require.NoError(t, err)
	require.NotNil(t, buf)
	require.Nil(t, err)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/auth/get-profile", buf)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	c.Set("uid", userReq.UID)

	mockAuthSvc.EXPECT().GetProfile(gomock.Any(), gomock.Eq(userReq)).Return(&dto.GetProfileRes{
		Status: 200,
		User:   user,
	}, nil)

	errRes := authController.GetProfile(c)

	require.NoError(t, errRes)
}
