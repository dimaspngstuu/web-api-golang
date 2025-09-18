package book

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]BooksModel, error)
	FindById(ID int) (BooksModel, error)
	Create(book BooksModel) (BooksModel, error)
	DeleteById(ID int, book BooksModel)
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

func (r *repository) DeleteById(ID int, book BooksModel) {
	var books BooksModel
	r.db.Where("ID = ?", ID).Delete(&books)

}
