package storage

import (
	"encoding/json"
	"os"

	"CARAPP-api/internal/models"
)

func LoadDeathComments(path string) ([]models.DeathComments, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var comments []models.DeathComments

	err = json.Unmarshal(data, &comments)
	if err != nil {
		return nil, err
	}

	return comments, nil
}
