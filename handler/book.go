package handler

import (
	"example/Go-Api/book"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewHandlerBook(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) GetBook(c *gin.Context) {
	StringId := c.Param("id")
	id, _ := strconv.Atoi(StringId)
	book, _ := h.bookService.FindById(id)

	if id <= 0 || id != book.ID {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID Tidak Valid",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})

}

func (h *bookHandler) GetListBooksHandler(c *gin.Context) {
	books, err := h.bookService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	var booksResponse []book.BooksResponse

	for _, b := range books {
		BooksResponse := book.BooksResponse{
			ID:          b.ID,
			Title:       b.Title,
			Description: b.Description,
			Price:       b.Price,
			Rating:      b.Rating,
		}
		booksResponse = append(booksResponse, BooksResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
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

func (h *bookHandler) UpdateBooksHandler(c *gin.Context) {
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

	IDString := c.Param("id")
	ID, _ := strconv.Atoi(IDString)

	book, err := h.bookService.Update(ID, bookRequest)

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

func (h *bookHandler) DeleteBooksHandler(c *gin.Context) {

	IDString := c.Param("id")
	ID, _ := strconv.Atoi(IDString)
	book, err := h.bookService.Delete(int(ID))

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
