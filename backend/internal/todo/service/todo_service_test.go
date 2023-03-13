package service

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/florentinuskev/simple-todo/internal/dao"
	"github.com/florentinuskev/simple-todo/internal/dto"
	mock_todo "github.com/florentinuskev/simple-todo/internal/todo/mock"
	"github.com/florentinuskev/simple-todo/public/utils"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type eqCreateTodoParamsMatcher struct {
	todo *dao.Todo
}

func (e eqCreateTodoParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(*dao.Todo)
	if !ok {
		return false
	}

	e.todo.ID = arg.ID

	return reflect.DeepEqual(e.todo, arg)
}

func (e eqCreateTodoParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v ", e.todo)
}

func EqCreateTodoParams(Todo *dao.Todo) gomock.Matcher {
	return eqCreateTodoParamsMatcher{Todo}
}

func TestGetTodos(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockTodoRepo := mock_todo.NewMockTodoRepository(ctrl)
	authSvc := NewTodoService(&utils.Config{}, mockTodoRepo)

	ctx := context.Background()

	uid := uuid.NewString()

	mockTodos := []*dao.Todo{{
		ID:   uuid.NewString(),
		UID:  uid,
		Todo: "Do Homework",
	}, {
		ID:   uuid.NewString(),
		UID:  uid,
		Todo: "Do Workout",
	}}

	mockTodoRepo.EXPECT().GetTodos(ctx, gomock.Eq(uid)).Return(mockTodos, nil)

	userRes, err := authSvc.GetTodos(ctx, &dto.GetTodosReq{UID: uid})

	require.NoError(t, err)
	require.NotNil(t, userRes)
	require.Equal(t, uint32(200), userRes.Status)
	require.NotNil(t, userRes.Todos)
}

func TestGetTodo(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockTodoRepo := mock_todo.NewMockTodoRepository(ctrl)
	authSvc := NewTodoService(&utils.Config{}, mockTodoRepo)

	ctx := context.Background()

	mockTodo := &dao.Todo{
		ID:   uuid.NewString(),
		UID:  uuid.NewString(),
		Todo: "Do Homework",
	}

	mockTodoRepo.EXPECT().GetTodo(ctx, gomock.Eq(mockTodo.ID)).Return(mockTodo, nil)

	userRes, err := authSvc.GetTodo(ctx, &dto.GetTodoReq{ID: mockTodo.ID})

	require.NoError(t, err)
	require.NotNil(t, userRes)
	require.Equal(t, uint32(200), userRes.Status)
	require.NotNil(t, userRes.Todo)
}

func TestNewTodo(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockTodoRepo := mock_todo.NewMockTodoRepository(ctrl)
	authSvc := NewTodoService(&utils.Config{}, mockTodoRepo)

	ctx := context.Background()

	mockTodo := &dao.Todo{
		UID:  uuid.NewString(),
		Todo: "Do Homework",
	}

	mockTodoRepo.EXPECT().NewTodo(ctx, EqCreateTodoParams(mockTodo)).Return(mockTodo, nil)

	userRes, err := authSvc.NewTodo(ctx, &dto.NewTodoReq{UID: mockTodo.UID, Todo: mockTodo.Todo})

	require.NoError(t, err)
	require.NotNil(t, userRes)
	require.Equal(t, uint32(201), userRes.Status)
	require.NotNil(t, userRes.Todo)
}

func TestEditTodo(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockTodoRepo := mock_todo.NewMockTodoRepository(ctrl)
	authSvc := NewTodoService(&utils.Config{}, mockTodoRepo)

	ctx := context.Background()

	mockTodo := &dao.Todo{
		ID:   uuid.NewString(),
		Todo: "Do Homework",
	}

	mockTodoRepo.EXPECT().EditTodo(ctx, gomock.Eq(mockTodo)).Return(mockTodo, nil)

	userRes, err := authSvc.EditTodo(ctx, &dto.EditTodoReq{ID: mockTodo.ID, Todo: mockTodo.Todo})

	require.NoError(t, err)
	require.NotNil(t, userRes)
	require.Equal(t, uint32(200), userRes.Status)
	require.NotNil(t, userRes.Todo)
}

func TestDeleteTodo(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockTodoRepo := mock_todo.NewMockTodoRepository(ctrl)
	authSvc := NewTodoService(&utils.Config{}, mockTodoRepo)

	ctx := context.Background()

	mockTodo := &dao.Todo{
		ID: uuid.NewString(),
	}

	mockTodoRepo.EXPECT().DeleteTodo(ctx, gomock.Eq(mockTodo)).Return(nil)

	userRes, err := authSvc.DeleteTodo(ctx, &dto.DeleteTodoReq{ID: mockTodo.ID})

	require.NoError(t, err)
	require.NotNil(t, userRes)
	require.Equal(t, uint32(200), userRes.Status)
	require.Equal(t, "Successfully delete the todo with given ID.", userRes.Msg)
}
