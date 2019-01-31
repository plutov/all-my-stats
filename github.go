package stats

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type githubUser struct {
	Followers int `json:"followers"`
}

func getGitHubUserFollowersCount() (int, error) {
	httpClient := &http.Client{
		Timeout: time.Second * 5,
	}
	resp, err := httpClient.Get(fmt.Sprintf("https://api.github.com/users/%s", os.Getenv("GH_USER")))
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	user := githubUser{}
	if err = json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return 0, err
	}

	return user.Followers, nil
}
