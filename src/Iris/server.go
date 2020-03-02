package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
)

// Author Model
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// Book Model
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

var books []Book
var mu sync.Mutex

func appendBooks() []Book {
	author := Author{
		FirstName: "James",
		LastName:  "Tan",
	}

	books := append(books, Book{
		ID:     "1",
		Isbn:   "44873",
		Title:  "Book One",
		Author: &author,
	})

	return books
}

func main() {

	app := iris.New()
	app.Use(recover.New())
	app.Post("/api/books/", func(ctx iris.Context) {
		fmt.Println("creating new book in go")
		var book Book
		ctx.ReadJSON(&book)
		book.ID = strconv.Itoa(rand.Intn(10000000))
		mu.Lock()
		books = append(books, book)
		mu.Unlock()
		ctx.JSON(book)
	})

	app.Run(iris.Addr(":8082"), iris.WithoutServerError(iris.ErrServerClosed))
}
