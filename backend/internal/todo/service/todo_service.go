package service

import (
	"context"

	"github.com/florentinuskev/simple-todo/internal/dao"
	"github.com/florentinuskev/simple-todo/internal/dto"
	"github.com/florentinuskev/simple-todo/internal/todo"
	"github.com/florentinuskev/simple-todo/public/utils"
	"github.com/google/uuid"
)

type TodoService struct {
	cfg *utils.Config
	r   todo.TodoRepository
}

func NewTodoService(cfg *utils.Config, r todo.TodoRepository) todo.TodoService {
	return &TodoService{cfg: cfg, r: r}
}

func (ts *TodoService) GetTodos(c context.Context, userReq *dto.GetTodosReq) (*dto.GetTodosRes, error) {
	todos, err := ts.r.GetTodos(c, userReq.UID)

	if err != nil {
		return nil, err
	}

	return &dto.GetTodosRes{
		Status: 200,
		Todos:  todos,
	}, nil
}
func (ts *TodoService) GetTodo(c context.Context, userReq *dto.GetTodoReq) (*dto.GetTodoRes, error) {
	todo, err := ts.r.GetTodo(c, userReq.ID)

	if err != nil {
		return nil, err
	}

	return &dto.GetTodoRes{
		Status: 200,
		Todo:   todo,
	}, nil
}
func (ts *TodoService) NewTodo(c context.Context, userReq *dto.NewTodoReq) (*dto.NewTodoRes, error) {
	todo := &dao.Todo{
		ID:   uuid.NewString(),
		UID:  userReq.UID,
		Todo: userReq.Todo,
	}

	newTodo, err := ts.r.NewTodo(c, todo)

	if err != nil {
		return nil, err
	}

	return &dto.NewTodoRes{
		Status: 201,
		Todo:   newTodo,
	}, nil
}
func (ts *TodoService) EditTodo(c context.Context, userReq *dto.EditTodoReq) (*dto.EditTodoRes, error) {
	editedTodo, err := ts.r.EditTodo(c, &dao.Todo{ID: userReq.ID, Todo: userReq.Todo})

	if err != nil {
		return nil, err
	}

	return &dto.EditTodoRes{
		Status: 200,
		Todo:   editedTodo,
	}, nil
}
func (ts *TodoService) DeleteTodo(c context.Context, userReq *dto.DeleteTodoReq) (*dto.DeleteTodoRes, error) {
	err := ts.r.DeleteTodo(c, &dao.Todo{ID: userReq.ID})

	if err != nil {
		return nil, err
	}

	return &dto.DeleteTodoRes{
		Status: 200,
		Msg:    "Successfully delete the todo with given ID.",
	}, nil
}
