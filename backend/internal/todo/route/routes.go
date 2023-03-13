package route

import (
	"github.com/florentinuskev/simple-todo/internal/middlewares"
	"github.com/florentinuskev/simple-todo/internal/todo"
	"github.com/florentinuskev/simple-todo/public/utils"
	"github.com/labstack/echo/v4"
)

func InitTodoRoute(e *echo.Echo, cfg *utils.Config, mw *middlewares.MiddlewareManager, ctrl todo.TodoController) {
	g := e.Group("/todos")

	g.Use(mw.IsAuthenticated)

	g.GET("/:uid", ctrl.GetTodos)
	g.GET("/get-todo/:id", ctrl.GetTodo)
	g.POST("/", ctrl.NewTodo)
	g.PATCH("/:id", ctrl.EditTodo)
	g.DELETE("/:id", ctrl.DeleteTodo)
}
