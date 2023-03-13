package todo

import (
	"context"

	"github.com/florentinuskev/simple-todo/internal/dto"
)

type TodoService interface {
	GetTodos(c context.Context, userReq *dto.GetTodosReq) (*dto.GetTodosRes, error)
	GetTodo(c context.Context, userReq *dto.GetTodoReq) (*dto.GetTodoRes, error)
	NewTodo(c context.Context, userReq *dto.NewTodoReq) (*dto.NewTodoRes, error)
	EditTodo(c context.Context, userReq *dto.EditTodoReq) (*dto.EditTodoRes, error)
	DeleteTodo(c context.Context, userReq *dto.DeleteTodoReq) (*dto.DeleteTodoRes, error)
}
