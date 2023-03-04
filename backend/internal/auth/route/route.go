package route

import (
	"github.com/florentinuskev/simple-todo/internal/auth"
	"github.com/florentinuskev/simple-todo/public/utils"
	"github.com/labstack/echo/v4"
)

type AuthRoute struct {
	e    *echo.Echo
	cfg  *utils.Config
	ctrl auth.AuthController
}

func InitAuthRoute(e *echo.Echo, ctrl auth.AuthController) {
	g := e.Group("/auth")

	g.POST("/register", ctrl.UserRegister)
	g.POST("/login", ctrl.UserLogin)
	g.GET("/get-profile", ctrl.GetProfile)
	g.GET("/refresh-token", ctrl.RefreshToken)
}
