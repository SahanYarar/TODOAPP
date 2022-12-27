package models

type UserResponse struct {
	ID         int    `json:"Id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	CreatedAt  string `json:"CreatedAt"`
	UpdatedAt  string `json:"UpdatedAt"`
	Deleted_at string `json:"DeletedAt"`
}
