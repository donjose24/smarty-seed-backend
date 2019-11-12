package unionbank

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jmramos02/smarty-seed-backend/config"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type AuthResponse struct {
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
	Error        string `json:"error"`
}

func GetAuthorizationCode(code string) (AuthResponse, error) {
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", config.GetUnionBankClientID())
	data.Set("redirect_uri", config.GetRedirectUri())
	data.Set("code", code)
	fmt.Println(data.Encode())

	request, err := http.NewRequest(http.MethodPost, config.GetUnionbankUrl()+"/customers/v1/oauth2/token", strings.NewReader(data.Encode()))

	request.Header.Add("accept", "application/json")
	request.Header.Add("content-type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		return AuthResponse{}, errors.New(fmt.Sprintf("Error getting access token: %v", err.Error))
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return AuthResponse{}, errors.New(fmt.Sprintf("Error parsing the response: %v", err.Error))
	}

	var response AuthResponse
	err = json.Unmarshal(body, &response)

	if err != nil {
		return AuthResponse{}, errors.New(fmt.Sprintf("Error reading the response: %v", err.Error))
	}

	return response, nil
}
