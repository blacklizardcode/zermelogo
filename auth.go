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

type AuthService struct {
	client *Client
}

func (c *Client) Auth() *AuthService {
	return &AuthService{c}
}

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

func (a *AuthService) SetBearerToken(bearer string) {
	a.client.apiKey = bearer
}
