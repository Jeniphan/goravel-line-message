package line_auth

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type LineLogin struct {
	ClientID       string
	ClientSecret   string
	RedirectURL    string
	AuthURL        string
	ProfileURL     string
	TokenURL       string
	RevokeURL      string
	VerifyTokenURL string
}

func NewLineLogin(clientID, clientSecret, redirectURL string) *LineLogin {
	return &LineLogin{
		ClientID:       clientID,
		ClientSecret:   clientSecret,
		RedirectURL:    redirectURL,
		AuthURL:        "https://access.line.me/oauth2/v2.1/authorize",
		ProfileURL:     "https://api.line.me/v2/profile",
		TokenURL:       "https://api.line.me/oauth2/v2.1/token",
		RevokeURL:      "https://api.line.me/oauth2/v2.1/revoke",
		VerifyTokenURL: "https://api.line.me/oauth2/v2.1/verify",
	}
}

func (l *LineLogin) GetLink(redirect_url string) string {
	link := fmt.Sprintf("%v?response_type=code&client_id=%v&redirect_uri=%v&scope=profile%%20openid%%20email&state=%v", l.AuthURL, l.ClientID, l.RedirectURL, redirect_url)
	return link
}

func (l *LineLogin) Refresh(token string) *TokenRefreshResponse {
	header := http.Header{}
	header.Set("Content-Type", "application/x-www-form-urlencoded")

	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", token)
	data.Set("client_id", l.ClientID)
	data.Set("client_secret", l.ClientSecret)

	response := sendHTTPRequest(l.TokenURL, "POST", header, strings.NewReader(data.Encode()))
	return parseTokenRefreshResponse(response)
}

func (l *LineLogin) Token(code, state string) *TokenResponse {
	header := http.Header{}
	header.Set("Content-Type", "application/x-www-form-urlencoded")

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", l.RedirectURL)
	data.Set("client_id", l.ClientID)
	data.Set("client_secret", l.ClientSecret)
	response := sendHTTPRequest(l.TokenURL, "POST", header, strings.NewReader(data.Encode()))
	// fmt.Println(response)
	return parseTokenResponse(response)
}

func (l *LineLogin) Profile(token string) *Profile {
	header := http.Header{}
	header.Set("Authorization", "Bearer "+token)
	response := sendHTTPRequest(l.ProfileURL, "GET", header, nil)
	var profile *Profile
	json.Unmarshal([]byte(response), &profile)
	return profile
}

func (l *LineLogin) Verify(token string) *TokenVerify {
	url := fmt.Sprintf("%v?access_token=%v", l.VerifyTokenURL, token)
	response := sendHTTPRequest(url, "GET", nil, nil)
	var tokenResponse *TokenVerify
	json.Unmarshal([]byte(response), &tokenResponse)
	return tokenResponse
}

func (l *LineLogin) Revoke(token string) string {
	header := http.Header{}
	header.Set("Content-Type", "application/x-www-form-urlencoded")

	data := url.Values{}
	data.Set("access_token", token)
	data.Set("client_id", l.ClientID)
	data.Set("client_secret", l.ClientSecret)

	return sendHTTPRequest(l.RevokeURL, "POST", header, strings.NewReader(data.Encode()))
}

func sendHTTPRequest(url, method string, header http.Header, body io.Reader) string {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Println("Failed to create request:", err)
	}
	req.Header = header
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request:", err)
	}
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
	}
	return string(responseBody)
}

func parseTokenResponse(response string) *TokenResponse {
	var tokenResponse TokenResponse
	json.Unmarshal([]byte(response), &tokenResponse)
	return &tokenResponse
}

func parseTokenRefreshResponse(response string) *TokenRefreshResponse {
	var tokenResponse TokenRefreshResponse
	json.Unmarshal([]byte(response), &tokenResponse)
	return &tokenResponse
}

func hashSHA256(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func getRemoteAddr() string {
	return "localhost"
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    uint   `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	IDToken      string `json:"id_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

type TokenRefreshResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    uint   `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

type TokenVerify struct {
	Scope     string `json:"scope"`
	ClientId  string `json:"client_id"`
	ExpiresIn uint   `json:"expires_in"`
}

type Profile struct {
	UserId        string `json:"userId"`
	DisplayName   string `json:"displayName"`
	PictureUrl    string `json:"pictureUrl"`
	StatusMessage string `json:"statusMessage"`
}
