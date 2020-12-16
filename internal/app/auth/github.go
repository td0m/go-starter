package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type githubData struct {
	Login string `json:"login"`
}

func getGithubData(accessToken string) (body *githubData, err error) {
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return
	}

	authorizationHeaderValue := fmt.Sprintf("token %s", accessToken)
	req.Header.Set("Authorization", authorizationHeaderValue)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&body)
	return
}
