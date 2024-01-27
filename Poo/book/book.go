package book

import (
	"fmt"
)

type Book struct {
	title  string
	author string
	pages  int
}

// simulando contructor
func NewBook(title string, author string, pages int) *Book {
	return &Book{
		title:  title,
		author: author,
		pages:  pages,
	}
}

// encapsulamiento
// set
func (b *Book) SetTitle(title string) {
	b.title = title
}

// Get
func (b *Book) Gettitle() string {
	return b.title
}

//----------------------------------------------------
//agregando metodos

func (b *Book) PrintInfo() {
	fmt.Printf("Title:%s\nAuthor:%s\nPages:%d\n", b.title, b.author, b.pages)
}

//----------------------------------------------------
//composicison -- Herencia

type TextBook struct {
	Book
	editorial string
	level     string
}

func NewTextBook(title string, author string, pages int, editorial string, level string) *TextBook {
	return &TextBook{
		Book: Book{
			title:  title,
			author: author,
			pages:  pages,
		},
		editorial: editorial,
		level:     level,
	}
}

func (b *TextBook) PrintInfo() {
	fmt.Printf("Title:%s\nAuthor:%s\nPages:%d\nEditorial:%s,Level:%s\n",
		b.title, b.author, b.pages, b.editorial, b.level)
}

//----------------------------------------------------
//Polimorfismo

type IPrintTable interface {
	PrintInfo()
}

func Print(p IPrintTable) {
	p.PrintInfo()
}
