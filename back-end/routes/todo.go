package routes

import (
	"todo-app/handlers"
	"todo-app/middleware"
	repositories "todo-app/repository"

	"todo-app/pkg/mysql"

	"github.com/labstack/echo/v4"
)

func TodoRoutes(e *echo.Group) {
	todoRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerTodo(todoRepository)

	e.POST("/todo", middleware.Auth(h.CreateTodo))

}
