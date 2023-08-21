package handlers

import (
	"encoding/json"
	"net/http"

	"groupie-tracker/backend/models"
)


func GetRelations(url string) (*models.Relations, error) {
    response, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()

    var relations models.Relations
    if err := json.NewDecoder(response.Body).Decode(&relations); err != nil {
        return nil, err
    }

    return &relations, nil
}
