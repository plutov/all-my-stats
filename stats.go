package stats

import (
	"html/template"
	"net/http"
)

type tplData struct {
	GitHubFollowers    int
	TwitterFollowers   int
	SOFReputation      int
	YouTubeSubscribers string
	YouTubeViews       string
	YouTubeTop5Videos  *youtubeVideos
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

	sofReputation, err := getSOFReputation()
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

	youtubeVideos, err := getYouTubeTop5Videos()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	gaStats, err := getGoogleAnalyticsStats()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	gaPages, err := getGoogleAnalyticsPages()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	data := tplData{
		GitHubFollowers:    githubFollowers,
		TwitterFollowers:   twitterFollowers,
		SOFReputation:      sofReputation,
		YouTubeSubscribers: youtubeStats.Items[0].Statistics.Subscribers,
		YouTubeViews:       youtubeStats.Items[0].Statistics.Views,
		YouTubeTop5Videos:  youtubeVideos,
		GAPages30Days:      make(map[string]string),
	}

	for _, report := range gaStats.Reports {
		rows := report.Data.Rows
		for _, row := range rows {
			data.GAUsers30Days = row.Metrics[0].Values[0]
		}
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
