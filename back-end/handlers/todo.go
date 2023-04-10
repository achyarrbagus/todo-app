package handlers

import (
	"fmt"
	"net/http"
	dto "todo-app/dto/result"
	tododto "todo-app/dto/todo"
	"todo-app/models"
	"todo-app/repository"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerTodo struct {
	TodoRepository repository.TodoRepository
}

func HandlerTodo(TodoRepository repository.TodoRepository) *handlerTodo {
	return &handlerTodo{TodoRepository}
}

func (h *handlerTodo) CreateTodo(c echo.Context) error {
	request := new(tododto.CreateTodoRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	// get jwt token with key "id"
	loginUser := c.Get("userLogin")
	isLoginUser := loginUser.(jwt.MapClaims)["ID"].(float64)
	fmt.Println("thu bulat")
	fmt.Println(isLoginUser)

	todo := models.Todo{
		Title:     c.FormValue("title"),
		Body:      c.FormValue("body"),
		StartDate: c.FormValue("startDate"),
		EndDate:   c.FormValue("endDate"),
		UserID:    int(isLoginUser),
	}
	data, err := h.TodoRepository.CreateTodo(todo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})

	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})

}
func convertTodo(u models.Todo) tododto.TodoResponse {
	return tododto.TodoResponse{
		Title:     u.Title,
		Body:      u.Body,
		StartDate: u.StartDate,
		EndDate:   u.EndDate,
	}
}
