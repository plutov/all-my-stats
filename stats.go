package stats

import (
	"html/template"
	"net/http"
)

type tplData struct {
	GitHubFollowers    int
	TwitterFollowers   int
	YouTubeSubscribers string
	YouTubeViews       string
	GAUsers30Days      string
	GAPages30Days      map[string]string
}

// Stats .
func Stats(w http.ResponseWriter, r *http.Request) {
	githubFollowers, err := getGitHubUserFollowersCount()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	twitterFollowers, err := getTwitterUserFollowersCount()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	youtubeStats, err := getYouTubeStats()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	data := tplData{
		GitHubFollowers:    githubFollowers,
		TwitterFollowers:   twitterFollowers,
		YouTubeSubscribers: youtubeStats.Items[0].Statistics.Subscribers,
		YouTubeViews:       youtubeStats.Items[0].Statistics.Views,
		GAPages30Days:      make(map[string]string),
	}

	gaStats, err := getGoogleAnalyticsStats()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	for _, report := range gaStats.Reports {
		rows := report.Data.Rows
		for _, row := range rows {
			data.GAUsers30Days = row.Metrics[0].Values[0]
		}
	}

	gaPages, err := getGoogleAnalyticsPages()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	for _, report := range gaPages.Reports {
		rows := report.Data.Rows
		for _, row := range rows {
			data.GAPages30Days[row.Dimensions[0]] = row.Metrics[0].Values[0]
		}
	}

	t := template.New("tpl")
	t, _ = t.Parse(tpl)

	t.Execute(w, data)
}
