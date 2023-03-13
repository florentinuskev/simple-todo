package todo

import (
	"context"

	"github.com/florentinuskev/simple-todo/internal/dao"
)

type TodoRepository interface {
	GetTodos(c context.Context, uid string) ([]*dao.Todo, error)
	GetTodo(c context.Context, id string) (*dao.Todo, error)
	NewTodo(c context.Context, todo *dao.Todo) (*dao.Todo, error)
	EditTodo(c context.Context, todo *dao.Todo) (*dao.Todo, error)
	DeleteTodo(c context.Context, todo *dao.Todo) error
}
