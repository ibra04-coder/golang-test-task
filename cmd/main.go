package main

import (
	"log"
	"net/http"

	"golang-test-task/internal/db"
	"golang-test-task/internal/handlers"
	"golang-test-task/internal/repository"
	"golang-test-task/internal/routes"
	"golang-test-task/internal/usecase"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	database, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	repo := repository.NewPostgresRepository(database)
	uc := usecase.NewService(repo)
	handler := handlers.NewNumberHandler(uc)
	router := routes.NewRouter(handler)

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}