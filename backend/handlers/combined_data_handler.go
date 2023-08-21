// backend/handlers/combined_data_handler.go
package handlers

import (
	"groupie-tracker/backend/models"
)

func GetArtistsWithRelations() (*models.CombinedData, error) {
	artists, err := GetArtists()
	if err != nil {
		return nil, err
	}

	relationsMap := make(map[int]*models.Relations)
	for i := range artists {
		relations, err := GetRelations(artists[i].Relations)
		if err != nil {
			return nil, err
		}
		relationsMap[artists[i].ID] = relations
	}

	combinedData := &models.CombinedData{
		Artists:       artists,
		RelationsData: relationsMap,
	}

	return combinedData, nil
}
