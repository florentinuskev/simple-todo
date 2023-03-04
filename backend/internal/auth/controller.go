package auth

import "github.com/labstack/echo/v4"

type AuthController interface {
	GetProfile(c echo.Context) error
	UserRegister(c echo.Context) error
	UserLogin(c echo.Context) error
	RefreshToken(c echo.Context) error
}
