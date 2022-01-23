package iterator

import (
	"fmt"
	"testing"
)

func TestBookShelf_Iterator(t *testing.T) {
	shelf := NewBookShelf()
	shelf.AddBook(&Book{Name: "Cloud Native programming with Golang"})
	shelf.AddBook(&Book{Name: "Design Data-Intensive Applications"})
	shelf.AddBook(&Book{Name: "Learn Data Structures and Algorithms with Golang"})
	it := shelf.Iterator()
	for it.HasNext() {
		fmt.Println(it.Next())
	}
}
