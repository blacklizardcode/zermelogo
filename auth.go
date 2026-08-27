package zermelogo

import "net/http"

type tokenResponse struct {
	Access_token string `json:"access_token"`
	Token_type string `json:"token_type"`
	Expires_in string `json:"expires_in"`
}

type AuthService struct {
	client *Client
}

func (c *Client) Auth() *AuthService {
	return &AuthService{c}
}

func (a *AuthService) CodeToBearer(code int) (error, string) {
	var Bearer tokenResponse
	err := a.client.do(http.MethodPost, "/api/oauth/token?code=" + string(code), nil, Bearer)
	if err != nil {
		return err, ""
	}
	return nil, Bearer.Access_token
}

func (a *AuthService) SetBearerToken(bearer string) {
	a.client.apiKey = bearer
	return
}