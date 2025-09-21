package book

type Service interface {
	FindAll() ([]BooksModel, error)
	FindById(ID int) (BooksModel, error)
	Create(book BookRequest) (BooksModel, error)
	Update(ID int, book BookRequest) (BooksModel, error)
	DeleteById(ID int) error
}

type service struct {
	repository Repository
}

// DeleteById implements Service.
func (s *service) DeleteById(ID int) error {
	panic("unimplemented")
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]BooksModel, error) {
	return s.repository.FindAll()
}

func (s *service) FindById(ID int) (BooksModel, error) {
	var book BooksModel
	book, err := s.repository.FindById(ID)
	return book, err
}

func (s *service) Create(bookRequest BookRequest) (BooksModel, error) {
	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()

	books := BooksModel{
		Title:       bookRequest.Title,
		Price:       int(price),
		Description: bookRequest.Description,
		Rating:      int(rating),
	}

	newBook, err := s.repository.Create(books)

	return newBook, err
}

func (s *service) Update(ID int, bookRequest BookRequest) (BooksModel, error) {
	book, err := s.repository.FindById(ID)

	if err != nil {
		return BooksModel{}, err
	}
	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()

	book.Title = bookRequest.Title
	book.Price = int(price)
	book.Rating = int(rating)
	book.Description = bookRequest.Description

	UpdateBook, err := s.repository.Update(ID, book)

	return UpdateBook, err
}
