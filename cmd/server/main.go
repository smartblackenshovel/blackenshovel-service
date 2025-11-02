package main

import (
	"log"
	"net/http"

	"blackenshovel-service/routes"
)

func main() {
	mux := routes.SetupRouter()

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
