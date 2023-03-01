package repository

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/florentinuskev/simple-todo/internal/dao"
	"github.com/florentinuskev/simple-todo/public/utils"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var u = &dao.User{
	ID:       uuid.NewString(),
	Username: "test123",
	Password: "test123",
}

var cfg = utils.Config{}

func TestCreateUser(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	defer sqlxDB.Close()

	repo := NewAuthRepo(&cfg, sqlxDB)

	rows := sqlmock.NewRows([]string{"id", "username", "password", "created_at", "updated_at"}).AddRow(u.ID, u.Username, u.Password, time.Now(), time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(CreateUserQuery)).WithArgs(u.ID, u.Username, u.Password).WillReturnRows(rows)

	user, err := repo.CreateUser(context.Background(), u)

	assert.NotNil(t, user)
	assert.NoError(t, err)
	require.Equal(t, user.ID, u.ID)
}

func TestFindUserById(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	defer sqlxDB.Close()

	repo := NewAuthRepo(&cfg, sqlxDB)

	rows := sqlmock.NewRows([]string{"id", "username", "password", "created_at", "updated_at"}).AddRow(u.ID, u.Username, u.Password, time.Now(), time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(FindUserByIdQuery)).WithArgs(u.ID).WillReturnRows(rows)

	user, err := repo.FindUserById(context.Background(), u.ID)

	assert.NotNil(t, user)
	assert.NoError(t, err)
}

func TestFindUserByUsername(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer mockDB.Close()

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	defer sqlxDB.Close()

	repo := NewAuthRepo(&cfg, sqlxDB)

	rows := sqlmock.NewRows([]string{"id", "username", "password", "created_at", "updated_at"}).AddRow(u.ID, u.Username, u.Password, time.Now(), time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(FindUserByUsernameQuery)).WithArgs(u.Username).WillReturnRows(rows)

	user, err := repo.FindUserByUsername(context.Background(), u.Username)

	assert.NotNil(t, user)
	assert.NoError(t, err)
	assert.Equal(t, u.ID, user.ID)
	assert.Equal(t, u.Username, user.Username)
}
