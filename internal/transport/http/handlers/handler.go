package handlers

import (
	"awesomeProject/internal/logger"
	"awesomeProject/internal/model"
	"awesomeProject/internal/service"
	"awesomeProject/internal/storage"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type UserHandler struct {
	storage        *storage.Storage
	userRepository model.UserRepository
	service        *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (s *UserHandler) CreateUser(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := s.userRepository.Create(user); err != nil {
		logger.Logger().Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	password := HashPassword(*user.Password)
	user.Password = &password
	return c.NoContent(http.StatusCreated)
}

func (s *UserHandler) GetUser(c echo.Context) error {
	username := c.Param("username")
	user, err := s.userRepository.Find(username)
	if err != nil {
		logger.Logger().Println(err)
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, user)
}

func (s *UserHandler) DeleteUser(c echo.Context) error {
	username := c.Param("username")
	if err := s.userRepository.Delete(username); err != nil {
		logger.Logger().Println(err)
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		logger.Logger().Println(err)
	}
	return string(bytes)
}
