package stats

var tpl = `GitHub Followers: {{.GitHubFollowers}}
Twitter Followers: {{.TwitterFollowers}}
YouTube Subscribers: {{.YouTubeSubscribers}}
YouTube Views: {{.YouTubeViews}}
pliutau.com users (30 days): {{.GAUsers30Days}}
pliutau.com top 5 pages (30 days):
{{ range $key, $val := .GAPages30Days }}{{ $val }} | {{ $key }}{{ end }}`
