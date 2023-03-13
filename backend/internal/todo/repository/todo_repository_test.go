package repository

import (
	"context"
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/florentinuskev/simple-todo/internal/dao"
	"github.com/florentinuskev/simple-todo/public/utils"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

var cfg = utils.Config{}

func TestGetTodos(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	defer sqlxDB.Close()

	repo := NewTodoRepository(&cfg, sqlxDB)

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

	rows := sqlmock.NewRows([]string{"id", "uid", "todo", "created_at", "updated_at"}).
		AddRow(mockTodos[0].ID, mockTodos[0].UID, mockTodos[0].Todo, time.Now(), time.Now()).
		AddRow(mockTodos[1].ID, mockTodos[1].UID, mockTodos[1].Todo, time.Now(), time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(GetTodosQuery)).WithArgs().WillReturnRows(rows)

	todos, err := repo.GetTodos(context.Background(), uid)

	require.NoError(t, err)
	require.NotNil(t, todos)
	require.Equal(t, mockTodos[0].UID, todos[0].UID)
}

func TestGetTodo(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	defer sqlxDB.Close()

	repo := NewTodoRepository(&cfg, sqlxDB)

	ctx := context.Background()

	todo := &dao.Todo{
		ID:   uuid.NewString(),
		UID:  uuid.NewString(),
		Todo: "Hello Todo!",
		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	rows := sqlmock.NewRows([]string{"id", "uid", "todo", "created_at", "updated_at"}).AddRow(todo.ID, todo.UID, todo.Todo, todo.CreatedAt, todo.UpdatedAt)

	mock.ExpectQuery(regexp.QuoteMeta(GetTodoQuery)).WithArgs(todo.ID).WillReturnRows(rows)

	res, err := repo.GetTodo(ctx, todo.ID)

	require.NoError(t, err)
	require.Nil(t, err)
	require.NotNil(t, res)
}

func TestNewTodo(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	defer sqlxDB.Close()

	repo := NewTodoRepository(&cfg, sqlxDB)

	ctx := context.Background()

	todo := &dao.Todo{
		ID:   uuid.NewString(),
		UID:  uuid.NewString(),
		Todo: "Hello Todo!",
		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	rows := sqlmock.NewRows([]string{"id", "uid", "todo", "created_at", "updated_at"}).AddRow(todo.ID, todo.UID, todo.Todo, todo.CreatedAt, todo.UpdatedAt)

	mock.ExpectQuery(regexp.QuoteMeta(NewTodoQuery)).WithArgs(todo.ID, todo.UID, todo.Todo).WillReturnRows(rows)

	res, err := repo.NewTodo(ctx, todo)

	require.NoError(t, err)
	require.Nil(t, err)
	require.NotNil(t, res)

}

func TestEditTodo(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	defer sqlxDB.Close()

	repo := NewTodoRepository(&cfg, sqlxDB)

	ctx := context.Background()

	todo := &dao.Todo{
		ID:   uuid.NewString(),
		UID:  uuid.NewString(),
		Todo: "Hello Todo!",
		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	rows := mock.NewRows([]string{"id", "uid", "todo", "created_at", "updated_at"}).AddRow(todo.ID, todo.UID, todo.Todo, todo.CreatedAt, todo.UpdatedAt)

	mock.ExpectQuery(regexp.QuoteMeta(UpdateTodoQuery)).WithArgs(todo.Todo, todo.ID).WillReturnRows(rows)

	res, err := repo.EditTodo(ctx, todo)

	require.NoError(t, err)
	require.Nil(t, err)
	require.NotNil(t, res)
}

func TestDeleteTodo(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	defer sqlxDB.Close()

	repo := NewTodoRepository(&cfg, sqlxDB)

	ctx := context.Background()

	todo := &dao.Todo{
		ID:   uuid.NewString(),
		UID:  uuid.NewString(),
		Todo: "Hello Todo!",
		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	mock.ExpectExec(regexp.QuoteMeta(DeleteTodoQuery)).WithArgs(todo.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	errRes := repo.DeleteTodo(ctx, todo)

	require.NoError(t, errRes)
	require.Nil(t, errRes)
}

func TestDeleteTodoError(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	defer sqlxDB.Close()

	repo := NewTodoRepository(&cfg, sqlxDB)

	ctx := context.Background()

	todo := &dao.Todo{
		ID:   uuid.NewString(),
		UID:  uuid.NewString(),
		Todo: "Hello Todo!",
		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	mock.ExpectExec(regexp.QuoteMeta(DeleteTodoQuery)).WithArgs(todo.ID).WillReturnResult(sqlmock.NewResult(0, 0))

	errRes := repo.DeleteTodo(ctx, todo)

	require.Error(t, errRes)
	require.NotNil(t, errRes)
}
