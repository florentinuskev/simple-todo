package server

import (
	"net/http"

	"github.com/florentinuskev/simple-todo/public/utils"
	"github.com/labstack/echo/v4"
)

type Server struct {
	cfg *utils.Config
	e   *echo.Echo
}

func NewServer(cfg *utils.Config) *Server {
	return &Server{
		cfg: cfg,
		e:   echo.New(),
	}
}

func (s *Server) RunServer() {

	s.e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	s.e.Logger.Fatal(s.e.Start(":" + s.cfg.Env["PORT"]))
}
