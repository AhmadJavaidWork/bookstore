package main

import (
	"net/http"
	"strconv"
)

func (cfg *Config) handleDeleteAuthor(w http.ResponseWriter, r *http.Request) {
	authorIDString := r.PathValue("id")
	authorID, err := strconv.ParseInt(authorIDString, 0, 32)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid author id", err)
		return
	}

	err = cfg.db.DeleteAuthor(r.Context(), int32(authorID))
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Author not found", err)
		return
	}

	respondWithJson(w, http.StatusOK, nil)
}
