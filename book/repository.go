package book

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]BooksModel, error)
	FindById(ID int) (BooksModel, error)
	Create(book BooksModel) (BooksModel, error)
	DeleteById(ID int) error
	Update(ID int, book BooksModel) (BooksModel, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]BooksModel, error) {
	var books []BooksModel
	err := r.db.Find(&books).Error
	return books, err
}

func (r *repository) FindById(ID int) (BooksModel, error) {
	var book BooksModel
	err := r.db.First(&book, ID).Error
	return book, err
}

func (r *repository) Create(book BooksModel) (BooksModel, error) {
	err := r.db.Create(&book).Error
	return book, err
}

func (r *repository) DeleteById(ID int) error {
	var book BooksModel

	err := r.db.First(&book, ID).Error
	if err != nil {
		fmt.Println("ID not found")
		return err
	}
	// Hapus data jika ditemukan
	if err := r.db.Delete(&book).Error; err != nil {
		fmt.Println("Failed to delete book")
		return err
	}

	// Jika berhasil
	return nil

}

func (r *repository) Update(ID int, book BooksModel) (BooksModel, error) {

	err := r.db.Save(&book).Error
	return book, err
}
