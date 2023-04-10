package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"
	dto "todo-app/dto/result"
	userdto "todo-app/dto/user"
	"todo-app/models"
	"todo-app/pkg/bcrypt"
	jwtToken "todo-app/pkg/jwt"
	"todo-app/repository"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerUser struct {
	UserRepository repository.UserRepository
}

func HandlerUser(UserRepository repository.UserRepository) *handlerUser {
	return &handlerUser{UserRepository}
}

func (h *handlerUser) Login(c echo.Context) error {

	request := new(userdto.CreateUserRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	user := models.User{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}

	// check user by email user
	dbuser, err := h.UserRepository.Login(user.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	isValid := bcrypt.CheckPasswordHash(user.Password, dbuser.Password)
	if !isValid {
		fmt.Println(isValid)
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "wrong email or password"})
	}

	//generate token
	claims := jwt.MapClaims{}
	claims["email"] = dbuser.Email
	claims["id"] = dbuser.ID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // 2 hours expired

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	loginResponse := userdto.UserResponse{
		Email:    dbuser.Email,
		Password: dbuser.Password,
		ID:       dbuser.ID,
		Token:    token,
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: loginResponse})

}

func (h *handlerUser) CreateUser(c echo.Context) error {

	request := new(userdto.CreateUserRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	requestPassword := c.FormValue("password")
	password, err := bcrypt.HashingPassword(requestPassword)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	user := models.User{
		Email:    c.FormValue("email"),
		Password: password,
	}

	newUser, err := h.UserRepository.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	data, _ := h.UserRepository.GetUser(newUser.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})

}
