package middlewares

import (
	"github.com/florentinuskev/simple-todo/internal/auth"
	"github.com/florentinuskev/simple-todo/public/utils"
)

type MiddlewareManager struct {
	cfg *utils.Config
	ar  auth.AuthRepository
}

func NewMiddlewareManager(cfg *utils.Config, ar auth.AuthRepository) *MiddlewareManager {
	return &MiddlewareManager{cfg: cfg, ar: ar}
}
