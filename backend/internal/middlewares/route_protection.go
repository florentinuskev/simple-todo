package middlewares

import (
	"net/http"

	"github.com/florentinuskev/simple-todo/public/utils"
	"github.com/labstack/echo/v4"
)

func (mw *MiddlewareManager) IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authentication")

		valid, user, err := utils.VerifyJWT(tokenString, mw.cfg.Env["JWT_SECRET"])

		if err != nil {
			return err
		}

		if !valid {
			c.JSON(http.StatusUnauthorized, "Unauthorized route path.")
		}

		c.Set("uid", user.ID)

		return next(c)
	}

}
