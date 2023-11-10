package models

import (
	"html"
	"strings"
)

type ToDoPatchRequest struct {
	Description string `json:"description,omitempty"`
	Status      string `json:"status,omitempty"`
}

func (u *ToDoPatchRequest) Validate() bool {
	isPayloadNil := true
	u.Description = html.EscapeString(strings.TrimSpace(u.Description))
	u.Status = html.EscapeString(strings.TrimSpace(u.Status))
	if u.Description == "" || u.Status == "" {

		return isPayloadNil
	}
	return !isPayloadNil

}
