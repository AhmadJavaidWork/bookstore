package main

import (
	"net/http"
	"strconv"
)

type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (cfg *Config) handleGetAllAuthors(w http.ResponseWriter, r *http.Request) {
	dbAuthors, err := cfg.db.GetAllAuthors(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve all authors", err)
		return
	}

	authors := []Author{}

	for _, dbAuthor := range dbAuthors {
		authors = append(authors, Author{
			ID:   int(dbAuthor.ID),
			Name: dbAuthor.Name,
		})
	}
	respondWithJson(w, http.StatusOK, authors)
}

func (cfg *Config) handleGetAuthor(w http.ResponseWriter, r *http.Request) {
	authorIDString := r.PathValue("id")
	authorID, err := strconv.ParseInt(authorIDString, 0, 32)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid author id", err)
		return
	}

	author, err := cfg.db.GetAuthorById(r.Context(), int32(authorID))
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Author not found", err)
		return
	}

	respondWithJson(w, http.StatusOK, Author{
		ID:   int(authorID),
		Name: author.Name,
	})
}
