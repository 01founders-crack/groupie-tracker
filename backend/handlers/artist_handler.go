package handlers

import (
	"encoding/json"
	"net/http"

	"groupie-tracker/backend/models"
)

func GetArtists() ([]models.Artist, error) {
	url := "https://groupietrackers.herokuapp.com/api/artists"

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var artists []models.Artist
	if err := json.NewDecoder(response.Body).Decode(&artists); err != nil {
		return nil, err
	}

	return artists, nil
}
