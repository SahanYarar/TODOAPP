package models

import (
	"html"
	"strings"
)

type UserSignRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (u *UserSignRequest) Validate() bool {

	isPayloadNil := true
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Password = html.EscapeString(strings.TrimSpace(u.Password))

	if u.Password == "" || u.Name == "" {

		return isPayloadNil
	}
	return !isPayloadNil

}
