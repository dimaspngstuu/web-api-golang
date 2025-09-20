package handler

import (
	"example/Go-Api/book"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewHandlerBook(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) GetRoot(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "dimas pangestu",
	})
}

func (h *bookHandler) GetHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}

func (h *bookHandler) BooksHandler(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *bookHandler) QueryHandler(c *gin.Context) {
	query := c.Query("title")
	c.JSON(http.StatusOK, gin.H{
		"title": query,
	})
}

func (h *bookHandler) AddBooksHandler(c *gin.Context) {
	var bookRequest book.BookRequest
	err := c.ShouldBindJSON(&bookRequest)

	if err != nil {
		errorMessage := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errors := fmt.Sprintf("Error on field %s, then condition %s", e.Field(), e.ActualTag())
			errorMessage = append(errorMessage, errors)

		}
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errorMessage,
		})
		return

	}

	book, err := h.bookService.Create(bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})

}
