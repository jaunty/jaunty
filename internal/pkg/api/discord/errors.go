package discord

import (
	"fmt"
	"strings"
)

// Error is an error returned by Discord's API.
type Error struct {
	Code    int                    `json:"code"`
	Errors  map[string]interface{} `json:"errors"`
	Message string                 `json:"message"`
}

// Error implements the error interface.
func (e *Error) Error() string {
	var sb strings.Builder
	sb.WriteString("discord: ")

	if e.Code > 0 {
		sb.WriteString(fmt.Sprintf("%d: ", e.Code))
	}

	if e.Message != "" {
		sb.WriteString(e.Message)
	}

	return sb.String()
}
