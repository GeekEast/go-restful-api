package main

// Find represents a method search a thing in books
func Find(books []Book, id string) int {
	for idx, book := range books {
		if book.ID == id {
			return idx
		}
	}
	return -1
}
