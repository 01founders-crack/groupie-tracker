// backend/models/combined_data.go
package models

type CombinedData struct {
	Artists       []Artist
	RelationsData map[int]*Relations // Map artist ID to relations
}
