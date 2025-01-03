package main

import (
	"encoding/json"
	"net/http"

	"github.com/ahmadjavaidwork/bookstore/backend/books/internal/database"
)

func (cfg *Config) handleCreateBook(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name     string `json:"name"`
		AuthorID int    `json:"author_id"`
	}
	decorder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decorder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	book, err := cfg.db.CreateBook(r.Context(), database.CreateBookParams{
		Name:     params.Name,
		AuthorID: int32(params.AuthorID),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create book", err)
		return
	}

	dbBook, err := cfg.db.GetBookById(r.Context(), book.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Book not found", err)
		return
	}

	respondWithJson(w, http.StatusCreated, Book{
		ID:         int(dbBook.ID),
		Name:       dbBook.Name,
		AuthorID:   int(dbBook.ID),
		AuthorName: dbBook.AuthorName,
	})
}
