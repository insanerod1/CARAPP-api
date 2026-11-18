package models

type AssistantComments struct {
	Event string `json:"event"`
	Comments []string `json:"comments"`
}