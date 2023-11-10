package models

import (
	"html"
	"strings"
)

type ToDoRequest struct {
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
	UserID      uint64 `json:"user_id" binding:"required"`
}

func (u *ToDoRequest) Validate() bool {
	isPayloadNil := true
	u.Description = html.EscapeString(strings.TrimSpace(u.Description))
	u.Status = html.EscapeString(strings.TrimSpace(u.Status))
	if u.Description == "" || u.Status == "" {

		return isPayloadNil
	}
	return !isPayloadNil

}
