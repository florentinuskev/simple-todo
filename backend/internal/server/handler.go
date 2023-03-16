package server

import (
	authRepository "github.com/florentinuskev/simple-todo/internal/auth/repository"
	"github.com/florentinuskev/simple-todo/internal/middlewares"
	todoRepository "github.com/florentinuskev/simple-todo/internal/todo/repository"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4/middleware"

	authService "github.com/florentinuskev/simple-todo/internal/auth/service"
	todoService "github.com/florentinuskev/simple-todo/internal/todo/service"

	authController "github.com/florentinuskev/simple-todo/internal/auth/controller"
	todoCotroller "github.com/florentinuskev/simple-todo/internal/todo/controller"

	authRoute "github.com/florentinuskev/simple-todo/internal/auth/route"
	todoRoute "github.com/florentinuskev/simple-todo/internal/todo/route"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func (s *Server) InitHandler() {
	// Initialize all repositories
	authRepo := authRepository.NewAuthRepository(s.cfg, s.db)
	todoRepo := todoRepository.NewTodoRepository(s.cfg, s.db)

	// Initialize all services
	authSvc := authService.NewAuthService(s.cfg, authRepo)
	todoSvc := todoService.NewTodoService(s.cfg, todoRepo)

	// Initialize all controllers
	authCtrl := authController.NewAuthController(s.cfg, authSvc)
	todoCtrl := todoCotroller.NewTodoController(s.cfg, todoSvc)

	// Initialize Middleware
	mw := middlewares.NewMiddlewareManager(s.cfg, authRepo)

	// Validator
	s.e.Validator = &CustomValidator{validator: validator.New()}
	s.e.Use(middleware.CORS())

	// Initialize all routes
	g := s.e.Group("/api/v1")

	authRoute.InitAuthRoute(g, s.cfg, mw, authCtrl)
	todoRoute.InitTodoRoute(g, s.cfg, mw, todoCtrl)

}
