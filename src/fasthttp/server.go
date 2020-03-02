package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
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
	router := fasthttprouter.New()
	router.POST("/api/books", func(ctx *fasthttp.RequestCtx) {
		fmt.Println("creating new book in go")
		var book Book
		json.Unmarshal(ctx.PostBody(), &book)
		book.ID = strconv.Itoa(rand.Intn(10000000))
		mu.Lock()
		books = append(books, book)
		mu.Unlock()
		json.NewEncoder(ctx).Encode(book)
	})

	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}
