package models

import (
	"html"
	"strings"
)

type UserPasswordRequest struct {
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

func (u *UserPasswordRequest) Validate() bool {

	u.Password = html.EscapeString(strings.TrimSpace(u.Password))
	isPasswordEqualAndNotNil := false
	if u.Password == u.ConfirmPassword && u.Password != "" {
		isPasswordEqualAndNotNil = true
		return isPasswordEqualAndNotNil
	}
	return isPasswordEqualAndNotNil
}
