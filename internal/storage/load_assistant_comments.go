package storage

import (
	"encoding/json"
	"os"

	"CARAPP-api/internal/models"
)

func LoadAssistantComments(path string) ([]models.AssistantComments, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var comments []models.AssistantComments

	err = json.Unmarshal(data, &comments)
	if err != nil {
		return nil, err
	}

	return comments, nil
}
