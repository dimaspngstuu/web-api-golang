package main

import (
	"example/Go-Api/book"
	"example/Go-Api/handler"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/pustaka_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("database not connected")
	}
	db.AutoMigrate(&book.BooksModel{})

	//=> FindALl()
	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewHandlerBook(bookService)

	bookRequest := book.BookRequest{
		Title: "A man who cant moved",
		Price: "200",
	}

	bookService.Create(bookRequest)
	// bookRepository.DeleteById(2)

	router := gin.Default()
	v1 := router.Group("/v1")

	v1.GET("/", bookHandler.GetRoot)

	v1.GET("/hello", bookHandler.GetHello)
	v1.GET("/books/:id", bookHandler.BooksHandler)
	v1.GET("/query", bookHandler.QueryHandler)
	v1.POST("/books", bookHandler.AddBooksHandler)
	router.Run()

}
