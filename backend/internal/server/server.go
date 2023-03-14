package server

import (
	"github.com/florentinuskev/simple-todo/public/utils"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type Server struct {
	cfg *utils.Config
	e   *echo.Echo
	db  *sqlx.DB
}

func NewServer(cfg *utils.Config, db *sqlx.DB) *Server {
	return &Server{
		cfg: cfg,
		e:   echo.New(),
		db:  db,
	}
}

func (s *Server) RunServer() {

	s.InitHandler()
	s.e.Logger.Fatal(s.e.Start(":" + s.cfg.Env["PORT"]))
}
