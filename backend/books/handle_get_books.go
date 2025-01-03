package main

import (
	"net/http"
	"strconv"
)

type Book struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	AuthorID   int    `json:"author_id"`
	AuthorName string `json:"author_name"`
}

func (cfg *Config) handleGetAllBooks(w http.ResponseWriter, r *http.Request) {
	dbBooks, err := cfg.db.GetAllBooks(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve all books", err)
		return
	}

	books := []Book{}

	for _, dbBook := range dbBooks {
		books = append(books, Book{
			ID:         int(dbBook.ID),
			Name:       dbBook.Name,
			AuthorID:   int(dbBook.AuthorID),
			AuthorName: dbBook.AuthorName,
		})
	}
	respondWithJson(w, http.StatusOK, books)
}

func (cfg *Config) handleGetBook(w http.ResponseWriter, r *http.Request) {
	bookIDString := r.PathValue("id")
	bookID, err := strconv.ParseInt(bookIDString, 0, 32)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid book id", err)
		return
	}

	book, err := cfg.db.GetBookById(r.Context(), int32(bookID))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Book not found", err)
		return
	}

	respondWithJson(w, http.StatusOK, Book{
		ID:         int(book.ID),
		Name:       book.Name,
		AuthorID:   int(book.AuthorID),
		AuthorName: book.AuthorName,
	})
}
