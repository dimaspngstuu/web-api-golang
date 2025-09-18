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
	// bookRepository := book.NewRepository(db)
	// books, _ := bookRepository.FindAll()
	// for _, v := range books {
	// 	fmt.Println("Price :", v.Price)
	// }

	// //=> FindById)
	// res, _ := bookRepository.FindById(1)
	// fmt.Println("Title :", res.Title)

	//=> Create
	// createNewBook := book.BooksModel{
	// 	ID:          2,
	// 	Title:       "you can to be heroes",
	// 	Description: "this book for man",
	// 	Rating:      10,
	// }

	// bookRepository.Create(createNewBook)
	book.Repository.DeleteById(book.BooksModel{}, 2)

	router := gin.Default()
	router.GET("/", handler.GetRoot)
	router.GET("/hello", handler.GetHello)
	router.GET("/books/:id", handler.BooksHandler)
	router.GET("/query", handler.QueryHandler)

	router.POST("/books", handler.AddBooksHandler)
	router.Run()

}
