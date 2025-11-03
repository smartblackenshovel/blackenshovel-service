package main

import (
	"log"
	"net/http"

	"blackenshovel-service/config"
	"blackenshovel-service/internal/database"
	"blackenshovel-service/routes"
)

func main() {
	config.Load()

	database.Connect()

	log.Println("Starting server on :8080")
	mux := routes.SetupRouter()

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
