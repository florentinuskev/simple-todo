package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/florentinuskev/simple-todo/internal/dao"
	"github.com/florentinuskev/simple-todo/internal/dto"
	mock_todo "github.com/florentinuskev/simple-todo/internal/todo/mock"
	"github.com/florentinuskev/simple-todo/public/utils"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestGetTodos(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthSvc := mock_todo.NewMockTodoService(ctrl)
	authController := NewTodoController(&utils.Config{}, mockAuthSvc)

	uid := uuid.NewString()

	userReq := &dto.GetTodosReq{UID: uid}

	buf, err := utils.AnyToBytesBuffer(userReq)

	require.NoError(t, err)
	require.NotNil(t, buf)
	require.Nil(t, err)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/todos/"+uid, buf)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	todos := []*dao.Todo{{
		ID:   uuid.NewString(),
		UID:  uid,
		Todo: "Hello Todo1",
	},
		{
			ID:   uuid.NewString(),
			UID:  uid,
			Todo: "Hello Todo2",
		},
	}

	mockAuthSvc.EXPECT().GetTodos(gomock.Any(), gomock.Eq(userReq)).Return(&dto.GetTodosRes{
		Status: 200,
		Todos:  todos,
	}, nil)

	errRes := authController.GetTodos(c)

	require.NoError(t, errRes)
}

func TestGetTodo(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthSvc := mock_todo.NewMockTodoService(ctrl)
	authController := NewTodoController(&utils.Config{}, mockAuthSvc)

	todo := &dao.Todo{
		ID:   uuid.NewString(),
		UID:  uuid.NewString(),
		Todo: "New Todo1",
	}

	userReq := &dto.GetTodoReq{ID: todo.ID}

	buf, err := utils.AnyToBytesBuffer(userReq)

	require.NoError(t, err)
	require.NotNil(t, buf)
	require.Nil(t, err)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/todos/get-todo/"+todo.ID, buf)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	mockAuthSvc.EXPECT().GetTodo(gomock.Any(), gomock.Eq(userReq)).Return(&dto.GetTodoRes{
		Status: 200,
		Todo:   todo,
	}, nil)

	errRes := authController.GetTodo(c)

	require.NoError(t, errRes)
}

func TestNewTodo(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthSvc := mock_todo.NewMockTodoService(ctrl)
	authController := NewTodoController(&utils.Config{}, mockAuthSvc)

	todo := &dao.Todo{
		ID:   uuid.NewString(),
		UID:  uuid.NewString(),
		Todo: "New Todo1",
	}

	userReq := &dto.NewTodoReq{UID: todo.UID, Todo: todo.Todo}

	buf, err := utils.AnyToBytesBuffer(userReq)

	require.NoError(t, err)
	require.NotNil(t, buf)
	require.Nil(t, err)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/todos/", buf)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	mockAuthSvc.EXPECT().NewTodo(gomock.Any(), gomock.Eq(userReq)).Return(&dto.NewTodoRes{
		Status: 200,
		Todo:   todo,
	}, nil)

	errRes := authController.NewTodo(c)

	require.NoError(t, errRes)
}

func TestEditTodo(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthSvc := mock_todo.NewMockTodoService(ctrl)
	authController := NewTodoController(&utils.Config{}, mockAuthSvc)

	todo := &dao.Todo{
		ID:   uuid.NewString(),
		UID:  uuid.NewString(),
		Todo: "New Todo1",
	}

	userReq := &dto.EditTodoReq{ID: todo.ID}

	buf, err := utils.AnyToBytesBuffer(userReq)

	require.NoError(t, err)
	require.NotNil(t, buf)
	require.Nil(t, err)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPatch, "/api/v1/todos/"+todo.ID, buf)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	mockAuthSvc.EXPECT().EditTodo(gomock.Any(), gomock.Eq(userReq)).Return(&dto.EditTodoRes{
		Status: 200,
		Todo:   todo,
	}, nil)

	errRes := authController.EditTodo(c)

	require.NoError(t, errRes)
}

func TestDeleteTodo(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthSvc := mock_todo.NewMockTodoService(ctrl)
	authController := NewTodoController(&utils.Config{}, mockAuthSvc)

	todo := &dao.Todo{
		ID:   uuid.NewString(),
		UID:  uuid.NewString(),
		Todo: "New Todo1",
	}

	userReq := &dto.DeleteTodoReq{ID: todo.ID}

	buf, err := utils.AnyToBytesBuffer(userReq)

	require.NoError(t, err)
	require.NotNil(t, buf)
	require.Nil(t, err)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/todos/"+todo.ID, buf)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	mockAuthSvc.EXPECT().DeleteTodo(gomock.Any(), gomock.Eq(userReq)).Return(&dto.DeleteTodoRes{
		Status: 200,
		Msg:    "Successfully delete the todo with given ID.",
	}, nil)

	errRes := authController.DeleteTodo(c)

	require.NoError(t, errRes)
}
