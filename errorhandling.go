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
		Message    string `json:"message"`
		Error      string `json:"error"`
		StatusCode int    `json:"statusCode"`
		Response   struct {
			Status  int    `json:"status"`
			Message string `json:"message"`
			Details string `json:"details"`
		} `json:"response"`
	}
	json.NewDecoder(resp.Body).Decode(&errResp)

	message := errResp.Message
	if message == "" {
		message = errResp.Error
	}
	if message == "" {
		message = errResp.Response.Message
	}
	if message == "" {
		message = errResp.Response.Details
	}

	statusCode := errResp.StatusCode
	if statusCode == 0 {
		statusCode = errResp.Response.Status
	}
	if statusCode == 0 {
		statusCode = resp.StatusCode
	}

	return &APIError{StatusCode: statusCode, Message: message}
}
