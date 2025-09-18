package main

import (
	"example/Go-Api/book"
	"example/Go-Api/handler"
	"fmt"
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

	books := book.BooksModel{}
	books.ID = 4
	books.Price = 20000
	books.Rating = 5
	books.Title = "dont afraid of the dark"
	books.Description = "dont you afraid the dark in your area"

	err = db.Create(&books).Error

	if err != nil {
		fmt.Println("=========================")
		fmt.Println("error creating a new book")
		fmt.Println("=========================")
	}

	var res []book.BooksModel
	err = db.First(&res).Error

	if err != nil {
		fmt.Println("=========================")
		fmt.Println("error put the first of book")
		fmt.Println("=========================")
	} else {
		for _, b := range res {
			fmt.Printf("Title :%s", b.Title)
		}
	}

	router := gin.Default()
	router.GET("/", handler.GetRoot)
	router.GET("/hello", handler.GetHello)
	router.GET("/books/:id", handler.BooksHandler)
	router.GET("/query", handler.QueryHandler)

	router.POST("/books", handler.AddBooksHandler)
	router.Run()

}
