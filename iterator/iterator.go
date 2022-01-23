package iterator

type Aggregate interface {
	Iterator() Iterator
}

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type Book struct {
	Name string
}

func (b Book) String() string {
	return b.Name
}

type BookShelf struct {
	books []*Book
	last  int
}

func NewBookShelf() *BookShelf {
	return &BookShelf{
		books: make([]*Book, 0),
		last:  0,
	}
}

func (s *BookShelf) GetBook(idx int) *Book {
	if idx >= len(s.books) {
		return nil
	}
	return s.books[idx]
}
func (s *BookShelf) AddBook(book *Book) {
	s.books = append(s.books, book)
	s.last++
}

func (s *BookShelf) Len() int {
	return s.last
}

func (s *BookShelf) Iterator() Iterator {
	return NewBookShelfIterator(s)
}

type BookShelfIterator struct {
	shelf *BookShelf
	index int
}

func NewBookShelfIterator(shelf *BookShelf) Iterator {
	return &BookShelfIterator{shelf: shelf}
}

func (s *BookShelfIterator) HasNext() bool {
	return s.index < s.shelf.Len()
}

func (s *BookShelfIterator) Next() interface{} {
	var book = s.shelf.GetBook(s.index)
	s.index++
	return book
}
