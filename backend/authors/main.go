package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ahmadjavaidwork/bookstore/backend/authors/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	port string
	db   *database.Queries
}

func main() {
	cfg := &Config{}
	err := setUpCfg(cfg)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    ":" + cfg.port,
		Handler: mux,
	}

	mux.HandleFunc("GET /", cfg.handleGetAllAuthors)
	mux.HandleFunc("GET /{id}", cfg.handleGetAuthor)
	mux.HandleFunc("POST /", cfg.handleCreateAuthor)
	mux.HandleFunc("DELETE /{id}", cfg.handleDeleteAuthor)

	log.Printf("Serving on port: %s", cfg.port)
	srv.ListenAndServe()
}

func setUpCfg(apiCfg *Config) error {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		return configError("PORT")
	}
	apiCfg.port = port

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		return configError("DB_USER")
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		return configError("DB_PASSWORD")
	}
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		return configError("DB_HOST")
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		return configError("DB_PORT")
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		return configError("DB_NAME")
	}

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	dbConn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return fmt.Errorf("error opening connection to database: %s", err)
	}

	apiCfg.db = database.New(dbConn)

	return nil
}

func configError(envVar string) error {
	return fmt.Errorf("%s environment variable is not set", envVar)
}
