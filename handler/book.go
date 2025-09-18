package handler

import (
	"example/Go-Api/book"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetRoot(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "dimas pangestu",
	})
}

func GetHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}

func BooksHandler(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func QueryHandler(c *gin.Context) {
	query := c.Query("title")
	c.JSON(http.StatusOK, gin.H{
		"title": query,
	})
}

func AddBooksHandler(c *gin.Context) {
	var addbook book.BooksModel
	err := c.ShouldBindJSON(&addbook)

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
	c.JSON(http.StatusOK, gin.H{
		"title": addbook.Title,
		"price": addbook.Price,
	})

}
