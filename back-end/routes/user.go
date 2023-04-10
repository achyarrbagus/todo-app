package routes

import (
	"todo-app/handlers"
	"todo-app/pkg/mysql"
	repositories "todo-app/repository"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	e.POST("/register", h.CreateUser)
	e.POST("/login", h.Login)
}
