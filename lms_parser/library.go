package lms_parser

import "fmt"     // For creating string buffer
import "strings" // For use trim function

// This is a default implementation of Book
// It does not need to be public because
// the public version of it doesn't used anywhere.
type Book struct {
	Name   string
	Author string
}

// This is a default implementation of Library
type Library struct {
	books []Book
	count []int
}

// methods book

// Make a new book by name and author
func NewBook(name, author string) *Book {
	return &Book{
		Name:   strings.TrimSpace(name),
		Author: strings.TrimSpace(author),
	}
}

// compare between two books.check if the
// book is same or not by name or author
// returns true if compare successful
func (b *Book) Compare(name, author string) bool {
	if b.Name == strings.TrimSpace(name) && b.Author == strings.TrimSpace(author) {
		return true
	}
	return false
}

// methods Library
// Make a new Library
func NewLibrary() Library {
	books := make([]Book, 0)
	c := make([]int, 0)
	return Library{
		books: books,
		count: c,
	}
}

// This is add books in library
// This takes books pointer for performance reason

func (lib *Library) Add(book *Book) {
	var index int = -1
	for i, b := range lib.books {
		if b.Compare(book.Name, book.Author) {
			index = i
			break
		}
	}
	if index == -1 {
		lib.books = append(lib.books, *book)
		lib.count = append(lib.count, 1)
	} else {
		lib.count[index] = lib.count[index] + 1
	}

}

// This function no need in implementation
// But When It need to access the Library struct's
// Book it can helpful.
func (lib *Library) Books() []Book {
	return lib.books
}

// This function return a Book struct from Books slice
// It only use to get books. Nothing else
// It is not performant at this time
// It doesn't remove the book struct only return a object
// and Prevent it to printing using inner methods
func (lib *Library) Get(name, author string) Book {
	var i int = -1
	for index, book := range lib.books {
		if book.Compare(name, author) {
			i = index
			break
		}
	}
	if i == -1 {
		return Book{Name: "None", Author: "None"}
	} else if lib.count[i] == 0 {
		return Book{Name: "None", Author: "None"}
	}
	lib.count[i] = lib.count[i] - 1
	return lib.books[i]
}

// printing Library

// This help to print library
// It is a book struct method because all Book object
// need this to make it string
func (b *Book) String() string {
	return fmt.Sprintf("'%v' by '%v'", b.Name, b.Author)
}

// This method prevent you to print the actual
// books slice. This is by default called by
// print methods in go language have.
func (lib Library) String() string {
	var buf string = "Books are:\n"
	var sep string = "\t"
	for i, b := range lib.books {
		if lib.count[i] != 0 {
			buf = buf + sep + b.String()
			sep = "\n\t"
		}
	}
	return buf
}
