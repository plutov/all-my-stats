package stats

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

type youtubeStats struct {
	Items []struct {
		Statistics struct {
			Subscribers string `json:"subscriberCount"`
			Views       string `json:"viewCount"`
		} `json:"statistics"`
	} `json:"items"`
}

func getYouTubeStats() (*youtubeStats, error) {
	httpClient := &http.Client{
		Timeout: time.Second * 5,
	}
	resp, err := httpClient.Get(fmt.Sprintf("https://www.googleapis.com/youtube/v3/channels?part=statistics&id=%s&key=%s", os.Getenv("YT_CHANNEL_ID"), os.Getenv("YT_API_KEY")))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	stats := new(youtubeStats)
	if err = json.NewDecoder(resp.Body).Decode(&stats); err != nil {
		return nil, err
	}

	if len(stats.Items) == 0 {
		return nil, errors.New("youtube channel not found")
	}

	return stats, nil
}
