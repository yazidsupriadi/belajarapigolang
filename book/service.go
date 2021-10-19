package book

type Service interface {
	FindAll() ([]Book, error)
	FindById(ID int) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
	Update(ID int, bookRequest BookRequest) (Book, error)
	Delete(ID int) (Book, error)
}
type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err

	//return s.repository.FindAll()
}
func (s *service) FindById(ID int) (Book, error) {
	book, err := s.repository.FindById(ID)
	return book, err
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {

	book := Book{
		Title:       bookRequest.Title,
		Price:       int(bookRequest.Price),
		Description: bookRequest.Description,
		Rating:      int(bookRequest.Rating),
	}
	newBook, err := s.repository.Create(book)
	return newBook, err
}
func (s *service) Update(ID int, bookRequest BookRequest) (Book, error) {

	book, err := s.repository.FindById(ID)

	book.Title = bookRequest.Title
	book.Price = int(bookRequest.Price)
	book.Description = bookRequest.Description
	book.Rating = int(bookRequest.Rating)

	newBook, err := s.repository.Update(book)
	return newBook, err
}

func (s *service) Delete(ID int) (Book, error) {

	book, err := s.repository.FindById(ID)

	newBook, err := s.repository.Delete(book)
	return newBook, err
}
