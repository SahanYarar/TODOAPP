package models

type ToDoPatchRequest struct {
	Description string `json:"description,omitempty"`
	Status      string `json:"status,omitempty"`
}
