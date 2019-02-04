package stats

var tpl = `GitHub Followers: {{.GitHubFollowers}}
Twitter Followers: {{.TwitterFollowers}}
Stackoverflow Reputation: {{.SOFReputation}}
YouTube Subscribers: {{.YouTubeSubscribers}}
YouTube Views: {{.YouTubeViews}}
YouTube Top 5 videos:
{{ range $key, $val := .YouTubeTop5Videos.Items }}{{ $val.Statistics.Views }} | {{ $val.Snippet.Title }}
{{ end }}pliutau.com users (30 days): {{.GAUsers30Days}}
pliutau.com top 5 pages (30 days):
{{ range $key, $val := .GAPages30Days }}{{ $val }} | {{ $key }}
{{ end }}`
