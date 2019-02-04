package stats

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
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

type youtubeSearch struct {
	Items []struct {
		ID struct {
			Kind    string `json:"kind"`
			VideoID string `json:"videoId"`
		} `json:"id"`
	} `json:"items"`
}

type youtubeVideos struct {
	Items []struct {
		Snippet struct {
			Title string `json:"title"`
		} `json:"snippet"`
		Statistics struct {
			Views string `json:"viewCount"`
		} `json:"statistics"`
	} `json:"items"`
}

func getYouTubeTop5Videos() (*youtubeVideos, error) {
	httpClient := &http.Client{
		Timeout: time.Second * 5,
	}
	resp, err := httpClient.Get(fmt.Sprintf("https://www.googleapis.com/youtube/v3/search?part=snippet,id&order=viewCount&maxResults=6&channelId=%s&key=%s", os.Getenv("YT_CHANNEL_ID"), os.Getenv("YT_API_KEY")))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	search := new(youtubeSearch)
	if err = json.NewDecoder(resp.Body).Decode(&search); err != nil {
		return nil, err
	}

	if len(search.Items) != 6 {
		return nil, errors.New("youtube videos not found")
	}

	var ids []string
	for _, i := range search.Items {
		if i.ID.Kind == "youtube#video" {
			ids = append(ids, i.ID.VideoID)
		}
	}

	respVideos, err := httpClient.Get(fmt.Sprintf("https://www.googleapis.com/youtube/v3/videos?part=statistics,snippet&id=%s&key=%s", strings.Join(ids, ","), os.Getenv("YT_API_KEY")))
	if err != nil {
		return nil, err
	}

	defer respVideos.Body.Close()

	videos := new(youtubeVideos)
	if err = json.NewDecoder(respVideos.Body).Decode(&videos); err != nil {
		return nil, err
	}

	return videos, nil
}
