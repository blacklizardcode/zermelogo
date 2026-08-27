package zermelogo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("api error (%d): %s", e.StatusCode, e.Message)
}

func parseAPIError(resp *http.Response) error {
	var errResp struct {
		Message string `json:"message"`
	}
	json.NewDecoder(resp.Body).Decode(&errResp)
	return &APIError{StatusCode: resp.StatusCode, Message: errResp.Message}
}