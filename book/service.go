package book

type Service interface {
	FindAll() ([]BooksModel, error)
	FindById(ID int) (BooksModel, error)
	Create(bookRequest BookRequest) (BooksModel, error)
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
	book := BooksModel{
		Title: bookRequest.Title,
		Price: int(price),
	}

	newBook, err := s.repository.Create(book)

	return newBook, err
}
