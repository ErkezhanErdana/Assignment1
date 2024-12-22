package main

import (
	"fmt"
)

type Book struct {
	ID         string
	Title      string
	Author     string
	IsBorrowed bool
}

type Library struct {
	Books map[string]Book
}

func (l *Library) AddBook(book Book) {
	l.Books[book.ID] = book
}

func (l *Library) BorrowBook(id string) {
	if book, exists := l.Books[id]; exists {
		if !book.IsBorrowed {
			book.IsBorrowed = true
			l.Books[id] = book
			fmt.Println("You have borrowed the book:", book.Title)
		} else {
			fmt.Println("The book is already borrowed.")
		}
	} else {
		fmt.Println("Book not found.")
	}
}

func (l *Library) ReturnBook(id string) {
	if book, exists := l.Books[id]; exists {
		if book.IsBorrowed {
			book.IsBorrowed = false
			l.Books[id] = book
			fmt.Println("You have returned the book:", book.Title)
		} else {
			fmt.Println("The book is not borrowed.")
		}
	} else {
		fmt.Println("Book not found.")
	}
}

func (l *Library) ListBooks() {
	fmt.Println("Available books:")
	for _, book := range l.Books {
		if !book.IsBorrowed {
			fmt.Printf("ID: %s, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
		}
	}
}

func main() {
	library := Library{Books: make(map[string]Book)}
	var choice int
	var id, title, author string

	for {
		fmt.Println("\n1. Add – Add a book to the library")
		fmt.Println("2. Borrow – Borrow a book")
		fmt.Println("3. Return – Return a borrowed book")
		fmt.Println("4. List – List available books")
		fmt.Println("5. Exit – Exit the program")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter book ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter book title: ")
			fmt.Scan(&title)
			fmt.Print("Enter book author: ")
			fmt.Scan(&author)
			book := Book{ID: id, Title: title, Author: author, IsBorrowed: false}
			library.AddBook(book)
			fmt.Println("Book added successfully.")
		case 2:
			fmt.Print("Enter book ID to borrow: ")
			fmt.Scan(&id)
			library.BorrowBook(id)
		case 3:
			fmt.Print("Enter book ID to return: ")
			fmt.Scan(&id)
			library.ReturnBook(id)
		case 4:
			library.ListBooks()
		case 5:
			fmt.Println("Exiting program.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
