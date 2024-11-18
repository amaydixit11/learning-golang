package routes

import (
	"net/http"
	"pokemon-api/handlers"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/pokemons", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			if id := r.URL.Query().Get("id"); id != "" {
				handlers.GetPokemon(w, r)
			} else {
				handlers.GetPokemons(w, r)
			}
		case http.MethodPost:
			handlers.CreatePokemon(w, r)
		case http.MethodPut:
			handlers.UpdatePokemon(w, r)
		case http.MethodDelete:
			handlers.DeletePokemon(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
