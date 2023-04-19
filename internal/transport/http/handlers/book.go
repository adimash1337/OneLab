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

// @Summary CreateBook
// @Security ApiKeyAuth
// @Tags book
// @Description create book
// @ID create-book
// @Accept  json
// @Produce  json
// @Param input body models.Book true "book info"
// @Success 201 {string} string "id"
// @Failure 400,404 {string} string "message"
// @Failure 500 {string} string "message"
// @Failure default {string} string "message"
// @Router /book/create [post]
func (h Manager) CreateBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.Book
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, fmt.Sprintf("Bind err:%s", err))
			return
		}
		var ValidateBook = validator.New()
		validationErr := ValidateBook.Struct(req)
		if validationErr != nil {
			logger.Logger().Println(validationErr)
			c.JSON(http.StatusBadRequest, fmt.Sprintf("Validation err:%s", validationErr))
			return
		}
		response, err := h.srv.Book.Create(c.Request.Context(), &req)
		if err != nil {
			logger.Logger().Println(err)
			c.JSON(http.StatusBadRequest, fmt.Sprintf("Create err:%s", err))
			return
		}

		c.JSON(http.StatusCreated, fmt.Sprintf("Book successfully created !!", response))
	}
}

// @Summary GetBookByAuthor
// @Tags book
// @Description get book by Author
// @ID get-book-by-author
// @Accept json
// @Produce json
// @Param author path string true "Author"
// @Success 200 {object} models.Book
// @Failure 400,404 {object} string "message"
// @Failure 500 {object} string "message"
// @Failure default {object} string "message"
// @Router /book/get_by_Author [get]
func (h Manager) GetBookByAuthor() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		queryParam := c.Query("author")
		if queryParam == "" {
			log.Println("query is empty")
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, "not found")
			c.Abort()
			return
		}
		get, err := h.srv.Book.GetByAuthor(ctx, queryParam)
		if err != nil {
			logger.Logger().Println(err)
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusFound, get)
	}
}

// @Summary GetBookByID
// @Tags book
// @Description get book by ID
// @ID get-book-by-id
// @Accept json
// @Produce json
// @Param book_id path string true "book id"
// @Success 200 {object} models.Book
// @Failure 400,404 {object} string "message"
// @Failure 500 {object} string "message"
// @Failure default {object} string "message"
// @Router /book/get_by_id [get]
func (h Manager) GetBookByID() gin.HandlerFunc {
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
		get, err := h.srv.Book.GetByID(ctx, queryParam)
		if err != nil {
			logger.Logger().Println(err)
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusFound, get)
	}
}

// @Summary DeleteBook
// @Security ApiKeyAuth
// @Tags book
// @Description delete book by author
// @ID delete-book
// @Accept json
// @Produce json
// @Param author path string true "author"
// @Success 200
// @Failure 400,404 {object} string "message"
// @Failure 500 {object} string "message"
// @Failure default {object} string "message"
// @Router /book/delete [get]
func (h Manager) DeleteBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		queryParam := c.Query("author")
		if queryParam == "" {
			log.Println("query is empty")
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, "not found")
			c.Abort()
			return
		}
		err := h.srv.Book.Delete(c.Request.Context(), queryParam)
		if err != nil {
			logger.Logger().Println(err)
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, "User is deleted")
	}
}

// @Summary ListBooks
// @Tags book
// @Description list books
// @ID list-books
// @Accept json
// @Produce json
// @Success 200 {object} []models.Book
// @Failure 400,404 {object} string "message"
// @Failure 500 {object} string "message"
// @Failure default {object} string "message"
// @Router /book/list [get]
func (h Manager) ListBooks() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := h.srv.Book.List(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}

		c.JSON(200, resp)
	}
}
