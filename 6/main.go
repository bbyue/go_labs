package main

import "fmt"

type Stringer interface {
	String() string
}

type Book struct {
	Title  string
	Author string
	ISBN   string
}

func (b Book) String() {
	fmt.Println("Книга: ", b.Title)
	fmt.Println("Автор: ", b.Author)
	fmt.Println("ISBN: ", b.ISBN)
}

func main() {
	book := Book{Title: "The Go Programming Language (1st edition)", Author: "Brian Kernighan", ISBN: "978-0-134-19055-6"}
	book.String()
}
