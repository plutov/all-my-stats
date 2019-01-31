package stats

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

type twitterUser struct {
	Followers int `json:"followers_count"`
}

func getTwitterUserFollowersCount() (int, error) {
	httpClient := &http.Client{
		Timeout: time.Second * 5,
	}
	resp, err := httpClient.Get(fmt.Sprintf("https://cdn.syndication.twimg.com/widgets/followbutton/info.json?screen_names=%s", os.Getenv("TWITTER_USER")))
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	users := []twitterUser{}
	if err = json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return 0, err
	}

	if len(users) == 0 {
		return 0, errors.New("twitter user not found")
	}

	return users[0].Followers, nil
}
