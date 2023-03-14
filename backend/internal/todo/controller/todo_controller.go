package controller

import (
	"log"
	"net/http"

	"github.com/florentinuskev/simple-todo/internal/dto"
	"github.com/florentinuskev/simple-todo/internal/todo"
	"github.com/florentinuskev/simple-todo/public/utils"
	"github.com/labstack/echo/v4"
)

type TodoController struct {
	cfg *utils.Config
	ts  todo.TodoService
}

func NewTodoController(cfg *utils.Config, ts todo.TodoService) todo.TodoController {
	return &TodoController{
		cfg: cfg,
		ts:  ts,
	}
}

func (tc *TodoController) GetTodos(c echo.Context) error {
	userReq := &dto.GetTodosReq{}
	c.Bind(userReq)

	sUID := c.Get("uid").(string)

	userReq.UID = sUID

	userRes, err := tc.ts.GetTodos(c.Request().Context(), userReq)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, userRes)
}

func (tc *TodoController) GetTodo(c echo.Context) error {
	userReq := &dto.GetTodoReq{}
	c.Bind(userReq)
	userReq.ID = c.Param("id")

	userRes, err := tc.ts.GetTodo(c.Request().Context(), userReq)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, userRes)
}

func (tc *TodoController) NewTodo(c echo.Context) error {
	userReq := &dto.NewTodoReq{}
	c.Bind(userReq)

	sUID := c.Get("uid").(string)

	userReq.UID = sUID

	userRes, err := tc.ts.NewTodo(c.Request().Context(), userReq)

	if err != nil {
		log.Println("ERRNO: ", err.Error())
		return err
	}

	return c.JSON(http.StatusCreated, userRes)
}

func (tc *TodoController) EditTodo(c echo.Context) error {
	userReq := &dto.EditTodoReq{}
	c.Bind(userReq)
	userReq.ID = c.Param("id")

	userRes, err := tc.ts.EditTodo(c.Request().Context(), userReq)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, userRes)
}

func (tc *TodoController) DeleteTodo(c echo.Context) error {
	userReq := &dto.DeleteTodoReq{}
	c.Bind(userReq)
	userReq.ID = c.Param("id")

	userRes, err := tc.ts.DeleteTodo(c.Request().Context(), userReq)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, userRes)
}
