package route

import (
	"github.com/florentinuskev/simple-todo/internal/auth"
	"github.com/florentinuskev/simple-todo/internal/middlewares"
	"github.com/florentinuskev/simple-todo/public/utils"
	"github.com/labstack/echo/v4"
)

func InitAuthRoute(e *echo.Group, cfg *utils.Config, mw *middlewares.MiddlewareManager, ctrl auth.AuthController) {
	g := e.Group("/auth")

	g.POST("/register", ctrl.UserRegister)
	g.POST("/login", ctrl.UserLogin)

	g.Use(mw.IsAuthenticated)
	g.GET("/profile", ctrl.GetProfile)
}
