package gemini

import "fmt"

// Error API Errors returned.
// @see https://docs.gemini.com/rest-api/#error-payload
type Error struct {
	Message string `json:"message"`
	Reason  string `json:"reason"`

	// StatusCode Direct response code from response.
	StatusCode int
}

func (e Error) Error() string {
	return fmt.Sprintf(
		"Gemini Error - http status: %d; reason: %s; error message: %s;",
		e.StatusCode,
		e.Reason,
		e.Message,
	)
}
