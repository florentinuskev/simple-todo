package controller

import (
	"errors"
	"log"
	"net/http"

	"github.com/florentinuskev/simple-todo/internal/auth"
	"github.com/florentinuskev/simple-todo/internal/dto"
	"github.com/florentinuskev/simple-todo/public/utils"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	cfg *utils.Config
	as  auth.AuthService
}

func NewAuthController(cfg *utils.Config, as auth.AuthService) auth.AuthController {
	return &AuthController{cfg: cfg, as: as}
}

func (ac *AuthController) GetProfile(c echo.Context) error {
	uid := c.Get("uid")
	log.Println("UID", uid)

	sUID := uid.(string)

	if uid == "" {
		return c.JSON(http.StatusBadRequest, "User ID not provided.")
	}

	res, err := ac.as.GetProfile(c.Request().Context(), &dto.GetProfileReq{UID: sUID})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (ac *AuthController) UserRegister(c echo.Context) error {
	userReq := &dto.UserRegisterReq{}
	c.Bind(userReq)

	res, err := ac.as.UserRegister(c.Request().Context(), userReq)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, res)
}

func (ac *AuthController) UserLogin(c echo.Context) error {
	userReq := &dto.UserLoginReq{}
	c.Bind(userReq)

	res, err := ac.as.UserLogin(c.Request().Context(), userReq)

	if err != nil {
		if err == errors.New("Username does not exists") {
			return c.JSON(http.StatusNotFound, "Username does not exists.")
		}

		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (ac *AuthController) RefreshToken(c echo.Context) error {
	return nil
}
