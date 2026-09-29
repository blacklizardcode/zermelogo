package zermelogo

import (
	"net/http"
	"net/url"
	"strconv"
)

type tokenResponse struct {
	Access_token string `json:"access_token"`
	Token_type   string `json:"token_type"`
	Expires_in   int    `json:"expires_in"`
}

// AuthService handles the OAuth exchange used by the Zermelo API.
type AuthService struct {
	client *Client
}

// Auth returns the authentication service for this client.
func (c *Client) Auth() *AuthService {
	return &AuthService{c}
}

// CodeToBearer exchanges a single use authentication code for a bearer token.
// Authentication codes can only be used once and expire quickly, so prefer
// reusing a token that you already have.
func (a *AuthService) CodeToBearer(code int) (string, error) {
	var Bearer tokenResponse
	path, err := url.Parse("/api/oauth/token")
	if err != nil {
		return "", err
	}
	q := path.Query()
	q.Add("code", strconv.Itoa(code))
	path.RawQuery = q.Encode()

	err = a.client.do(http.MethodPost, path.String(), nil, &Bearer)
	if err != nil {
		return "", err
	}
	return Bearer.Access_token, nil
}

// SetBearerToken stores a bearer token that the client sends with every request.
func (a *AuthService) SetBearerToken(bearer string) {
	a.client.apiKey = bearer
}
