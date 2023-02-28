package auth

import "github.com/labstack/echo/v4"

type AuthController interface {
	GetUser(c echo.Context) error
	UserRegister(c echo.Context) error
	UserLogin(c echo.Context) error
	RefreshToken(c echo.Context) error
}
