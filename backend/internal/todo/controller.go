package todo

import "github.com/labstack/echo/v4"

type TodoController interface {
	GetTodos(c echo.Context) error
	GetTodo(c echo.Context) error
	NewTodo(c echo.Context) error
	EditTodo(c echo.Context) error
	DeleteTodo(c echo.Context) error
}
