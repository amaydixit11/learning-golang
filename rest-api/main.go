package main

import (
	"log"
	"net/http"
	"pokemon-api/routes"
)

func main() {
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)

	log.Println("Starting Pokémon API server on :1755...")
	log.Fatal(http.ListenAndServe(":1755", mux))
}
