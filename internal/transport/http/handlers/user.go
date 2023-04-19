package handlers

import (
	"awesomeProject/internal/logger"
	"awesomeProject/internal/models"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"time"
)

// @Summary CreateUser
// @Tags auth
// @Description create user
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body models.User true "account info"
// @Success 200 {string} string "id"
// @Failure 400,404 {string} string "message"
// @Failure 500 {string} string "message"
// @Failure default {string} string "message"
// @Router /user/create [post]
func (h Manager) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.User
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, fmt.Sprintf("Bind err:%s", err))
			return
		}
		var Validate = validator.New()
		validationErr := Validate.Struct(req)
		if validationErr != nil {
			logger.Logger().Println(validationErr)
			c.JSON(http.StatusBadRequest, fmt.Sprintf("Validation err:%s", validationErr))
			return
		}
		response, err := h.srv.User.Create(c.Request.Context(), req)
		if err != nil {
			logger.Logger().Println(err)
			c.JSON(http.StatusBadRequest, fmt.Sprintf("Create err:%s", err))
			return
		}

		c.JSON(http.StatusCreated, fmt.Sprintf("Successfully Signed Up!!", response))
	}
}

// @Summary Login
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body models.UserAuth true "credentials"
// @Success 200 {string} string "ID, Name"
// @Failure 400,404 {string} string "message"
// @Failure 500 {string} string "message"
// @Failure default {string} string "message"
// @Router /user/login [post]
func (h Manager) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.UserAuth
		if err := c.Bind(&req); err != nil {
			c.JSON(http.StatusBadRequest, fmt.Sprintf("Bind err:%s", err))
		}

		userData, err := h.srv.User.Login(c.Request.Context(), &req)
		if err != nil {
			logger.Logger().Println(err)
			c.JSON(http.StatusBadRequest, fmt.Sprintf("Login err:%s", err))
			return
		}

		c.JSON(302, userData)
	}
}

// @Summary UpdatePassword
// @Security ApiKeyAuth
// @Tags auth
// @Description update password
// @ID update-password
// @Accept json
// @Produce json
// @Param input body models.UpdatePasswordReq true "current and new password"
// @Success 200
// @Failure 400,404 {string} string "message"
// @Failure 500 {string} string "message"
// @Failure default {string} string "message"
// @Router /update-pass [put]
func (h Manager) UpdatePassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.UpdatePasswordReq
		if err := c.Bind(&req); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}

		queryParam := c.Query("username")
		if queryParam == "" {
			log.Println("query is empty")
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, "not found")
			c.Abort()
			return
		}
		if err := h.srv.User.UpdatePassword(c.Request.Context(), &req, queryParam); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}

		c.JSON(http.StatusOK, req)
	}
}

// @Summary GetUserByUsername
// @Tags user
// @Description get users by username
// @ID get-users-by-username
// @Accept json
// @Produce json
// @Success 200 {object} []models.User
// @Failure 400,404 {string} string "message"
// @Failure 500 {string} string "message"
// @Failure default {string} string "message"
// @Router /user/get_by_name [get]
func (h Manager) GetUserByUsername() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		queryParam := c.Query("username")
		if queryParam == "" {
			log.Println("query is empty")
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, "not found")
			c.Abort()
			return
		}
		get, err := h.srv.User.GetByUsername(ctx, queryParam)
		if err != nil {
			logger.Logger().Println(err)
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusFound, get)
	}
}

// @Summary GetUserByID
// @Tags user
// @Description get users by id
// @ID get-users-by_id
// @Accept json
// @Produce json
// @Success 200 {object} []models.User
// @Failure 400,404 {string} string "message"
// @Failure 500 {string} string "message"
// @Failure default {string} string "message"
// @Router /user/get_by_id [get]
func (h Manager) GetUserByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		queryParam := c.Query("id")
		if queryParam == "" {
			log.Println("query is empty")
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, "not found")
			c.Abort()
			return
		}
		get, err := h.srv.User.GetByID(ctx, queryParam)
		if err != nil {
			logger.Logger().Println(err)
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusFound, get)
	}
}

// @Summary DeleteUser
// @Tags user
// @Description delete user by username
// @ID delete-user-by-username
// @Accept json
// @Produce json
// @Success 200 {string} string "message"
// @Failure 400,404 {string} string "message"
// @Failure 500 {string} string "message"
// @Failure default {string} string "message"
// @Router /user/delete [get]
func (h Manager) DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		queryParam := c.Query("username")
		if queryParam == "" {
			log.Println("query is empty")
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, "not found")
			c.Abort()
			return
		}
		err := h.srv.User.Delete(c.Request.Context(), queryParam)
		if err != nil {
			logger.Logger().Println(err)
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, "User is deleted")
	}
}
