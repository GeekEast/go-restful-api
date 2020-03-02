package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
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

func getBooks(res http.ResponseWriter, _ *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(books) // slice of struct to json obj
}

func getBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	id := params["id"]
	idx := Find(books, id)
	book := books[idx]
	json.NewEncoder(res).Encode(book)
}

func createBook(res http.ResponseWriter, req *http.Request) {
	fmt.Println("creating new book")
	res.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(req.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	mu.Lock()
	books = append(books, book)
	mu.Unlock()
	json.NewEncoder(res).Encode(book)
}

func updateBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	id := params["id"]
	idx := Find(books, id)
	if idx != -1 {

		var nextBook Book
		_ = json.NewDecoder(req.Body).Decode(&nextBook)
		nextBook.ID = id

		mu.Lock()
		books = append(books[:idx], books[idx+1:]...)
		books = append(books, nextBook)
		mu.Unlock()

		json.NewEncoder(res).Encode(nextBook)
	}
}

func deleteBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	id := params["id"]
	idx := Find(books, id)
	if idx != -1 {
		mu.Lock()
		books = append(books[:idx], books[idx+1:]...)
		mu.Unlock()
	}
	json.NewEncoder(res).Encode(books)
}

func main() {
	router := mux.NewRouter()
	books = appendBooks()
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books/", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
