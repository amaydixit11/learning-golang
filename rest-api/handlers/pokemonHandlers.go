package handlers

import (
	"encoding/json"
	"net/http"

	"pokemon-api/models"
)

var pokemons = []models.Pokemon{}

func GetPokemons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemons)
}

func GetPokemon(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	for _, pokemon := range pokemons {
		if pokemon.PokedexNumber == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(pokemon)
			return
		}
	}
	http.Error(w, "Pokémon not found", http.StatusNotFound)
}

func CreatePokemon(w http.ResponseWriter, r *http.Request) {
	var pokemon models.Pokemon
	if err := json.NewDecoder(r.Body).Decode(&pokemon); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	pokemons = append(pokemons, pokemon)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemon)
}

func UpdatePokemon(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var updatedPokemon models.Pokemon
	if err := json.NewDecoder(r.Body).Decode(&updatedPokemon); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	for i, pokemon := range pokemons {
		if pokemon.PokedexNumber == id {
			updatedPokemon.PokedexNumber = id
			pokemons[i] = updatedPokemon
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedPokemon)
			return
		}
	}

	http.Error(w, "Pokémon not found", http.StatusNotFound)
}

func DeletePokemon(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	for i, pokemon := range pokemons {
		if pokemon.PokedexNumber == id {
			pokemons = append(pokemons[:i], pokemons[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Pokémon not found", http.StatusNotFound)
}
