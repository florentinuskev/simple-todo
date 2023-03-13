package server

import (
	authRepository "github.com/florentinuskev/simple-todo/internal/auth/repository"
	"github.com/florentinuskev/simple-todo/internal/middlewares"
	todoRepository "github.com/florentinuskev/simple-todo/internal/todo/repository"

	authService "github.com/florentinuskev/simple-todo/internal/auth/service"
	todoService "github.com/florentinuskev/simple-todo/internal/todo/service"

	authController "github.com/florentinuskev/simple-todo/internal/auth/controller"
	todoCotroller "github.com/florentinuskev/simple-todo/internal/todo/controller"

	authRoute "github.com/florentinuskev/simple-todo/internal/auth/route"
	todoRoute "github.com/florentinuskev/simple-todo/internal/todo/route"
)

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

	// Initialize all routes
	authRoute.InitAuthRoute(s.e, s.cfg, mw, authCtrl)
	todoRoute.InitTodoRoute(s.e, s.cfg, mw, todoCtrl)

}
