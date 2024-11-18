package models

type Pokemon struct {
	PokedexNumber string  `json:"pokedexnumber"` // Unique identifier
	Name          string  `json:"name"`          // Name of the Pokémon
	PrimaryType   string  `json:"primarytype"`   // Types like Fire, Water, etc.
	SecondaryType *string `json:"secondarytype"` // Types like Fire, Water, etc.
}
