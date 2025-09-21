package book

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]BooksModel, error)
	FindById(ID int) (BooksModel, error)
	Create(book BooksModel) (BooksModel, error)
	Update(ID int, book BooksModel) (BooksModel, error)
	Delete(book BooksModel) (BooksModel, error)
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

func (r *repository) Update(ID int, book BooksModel) (BooksModel, error) {

	err := r.db.Save(&book).Error
	return book, err
}

func (r *repository) Delete(book BooksModel) (BooksModel, error) {
	err := r.db.Delete(&book).Error
	return book, err
}
