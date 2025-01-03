package main

import (
	"encoding/json"
	"net/http"
)

func (cfg *Config) handleCreateAuthor(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decorder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decorder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	author, err := cfg.db.CreateAuthor(r.Context(), params.Name)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create author", err)
		return
	}

	respondWithJson(w, http.StatusCreated, Author{
		ID:   int(author.ID),
		Name: author.Name,
	})
}
