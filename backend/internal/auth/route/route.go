package route

import (
	"github.com/florentinuskev/simple-todo/internal/auth"
	"github.com/florentinuskev/simple-todo/internal/middlewares"
	"github.com/labstack/echo/v4"
)

func InitAuthRoute(e *echo.Echo, ctrl auth.AuthController, mw *middlewares.MiddlewareManager) {
	g := e.Group("/auth")

	g.POST("/register", ctrl.UserRegister)
	g.POST("/login", ctrl.UserLogin)

	g.Use(mw.IsAuthenticated)
	g.GET("/profile/", ctrl.GetProfile)
}
