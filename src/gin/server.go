package main

import (
	"math/rand"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
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
	router := gin.Default()
	router.POST("/api/books", func(ctx *gin.Context) {
		var book Book
		ctx.Bind(&book)
		book.ID = strconv.Itoa(rand.Intn(10000000))
		mu.Lock()
		books = append(books, book)
		mu.Unlock()
		ctx.JSON(200, book)
	})

	router.Run(":8080")

}
