package repository

import (
	"context"

	"github.com/florentinuskev/simple-todo/internal/auth"
	"github.com/florentinuskev/simple-todo/internal/dao"
	"github.com/florentinuskev/simple-todo/public/utils"
	"github.com/jmoiron/sqlx"
)

type AuthRepo struct {
	cfg *utils.Config
	db  *sqlx.DB
}

func NewAuthRepository(cfg *utils.Config, db *sqlx.DB) auth.AuthRepository {
	return &AuthRepo{cfg: cfg, db: db}
}

func (ar *AuthRepo) CreateUser(c context.Context, user *dao.User) (*dao.User, error) {
	u := &dao.User{}

	if err := ar.db.QueryRowxContext(c, CreateUserQuery, user.ID, user.Username, user.Password).StructScan(u); err != nil {
		return nil, err
	}

	return u, nil
}

func (ar *AuthRepo) FindUserById(c context.Context, uid string) (*dao.User, error) {
	u := &dao.User{}

	if err := ar.db.QueryRowxContext(c, FindUserByIdQuery, uid).StructScan(u); err != nil {
		return nil, err
	}

	return u, nil
}

func (ar *AuthRepo) FindUserByUsername(c context.Context, username string) (*dao.User, error) {
	u := &dao.User{}

	if err := ar.db.QueryRowxContext(c, FindUserByUsernameQuery, username).StructScan(u); err != nil {
		return nil, err
	}

	return u, nil
}
