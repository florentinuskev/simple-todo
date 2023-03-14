package repository

import (
	"context"
	"errors"

	"github.com/florentinuskev/simple-todo/internal/dao"
	"github.com/florentinuskev/simple-todo/internal/todo"
	"github.com/florentinuskev/simple-todo/public/utils"
	"github.com/jmoiron/sqlx"
)

type TodoRepository struct {
	cfg *utils.Config
	db  *sqlx.DB
}

func NewTodoRepository(cfg *utils.Config, db *sqlx.DB) todo.TodoRepository {
	return &TodoRepository{cfg: cfg, db: db}
}

func (tr *TodoRepository) GetTodos(c context.Context, uid string) ([]*dao.Todo, error) {

	var todos []*dao.Todo

	rows, err := tr.db.QueryxContext(c, GetTodosQuery, uid)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		t := &dao.Todo{}

		if err := rows.StructScan(t); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}

	return todos, nil
}

func (tr *TodoRepository) GetTodo(c context.Context, id string) (*dao.Todo, error) {
	t := &dao.Todo{}

	if err := tr.db.QueryRowxContext(c, GetTodoQuery, id).StructScan(t); err != nil {
		return nil, err
	}

	return t, nil
}

func (tr *TodoRepository) NewTodo(c context.Context, todo *dao.Todo) (*dao.Todo, error) {
	t := &dao.Todo{}

	if err := tr.db.QueryRowxContext(c, NewTodoQuery, todo.UID, todo.Todo).StructScan(t); err != nil {
		return nil, err
	}

	return t, nil
}

func (tr *TodoRepository) EditTodo(c context.Context, todo *dao.Todo) (*dao.Todo, error) {
	t := &dao.Todo{}

	if err := tr.db.QueryRowxContext(c, UpdateTodoQuery, todo.Todo, todo.ID).StructScan(t); err != nil {
		return nil, err
	}

	return t, nil
}

func (tr *TodoRepository) DeleteTodo(c context.Context, todo *dao.Todo) error {

	res, err := tr.db.ExecContext(c, DeleteTodoQuery, todo.ID)

	if err != nil {
		return err
	}

	affectedRow, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if affectedRow == 0 {
		return errors.New("NoAffectedRows")
	}

	return nil
}
