package main

import (
	"net/http"
	"strconv"
)

func (cfg *Config) handleDeleteBook(w http.ResponseWriter, r *http.Request) {
	bookIDString := r.PathValue("id")
	bookID, err := strconv.ParseInt(bookIDString, 0, 32)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid book id", err)
		return
	}

	err = cfg.db.DeleteBook(r.Context(), int32(bookID))
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Book not found", err)
		return
	}

	respondWithJson(w, http.StatusOK, nil)
}
