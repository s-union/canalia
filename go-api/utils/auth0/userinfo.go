package auth0

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type UserInfo struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
}

func FetchUserInfo(token string) (*UserInfo, error) {
	domain := os.Getenv("AUTH0_DOMAIN")
	url := "https://" + domain + "/userinfo"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch user info: status code %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)

	var userInfo UserInfo
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		return nil, err
	}

	return &userInfo, nil
}
