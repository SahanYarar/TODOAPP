package models

type UserResetPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}
