# all-my-stats

Google Cloud Function to display various social stats (pretty much specific for me).

- Google Analytics: users and top 5 pages (last 30 days)
- YouTube channel: subscribers count and top 5 videos (all time)
- GitHub: followers count
- Twitter: followers count
- Stackoverflow: reputation

## Requirements

1. Enable [Google Analytics Reporting API](https://console.developers.google.com/apis/library/analyticsreporting.googleapis.com).
2. Copy GCF service account email and add it in Google Analytics with Read & Analyze permission.
3. Enable [YouTube Data API](https://console.cloud.google.com/apis/api/youtube.googleapis.com)
4. Create YouTube API Key.

p.s. We can get GitHub / Twitter followers count without API keys.

## Deploy

```bash
YT_API_KEY= ./deploy.sh
```

## Response

[https://europe-west1-alexsandbox.cloudfunctions.net/stats](https://europe-west1-alexsandbox.cloudfunctions.net/stats)

```
GitHub Followers: 188
Twitter Followers: 364
Stackoverflow Reputation: 14275
YouTube Subscribers: 2457
YouTube Views: 44332
YouTube Top 5 videos:
8145 | packagemain #0: Building gRPC blockchain Server &amp; Client in Go
4730 | packagemain #12: Microservices with go-kit. Part 1
4551 | packagemain #4: Image recognition in Go using TensorFlow
4240 | packagemain #11: Getting started with OAuth2 in Go
4226 | packagemain #5: Face Detection in Go using OpenCV and MachineBox
pliutau.com users (30 days): 2353
pliutau.com top 5 pages (30 days):
174 | /
232 | /google-cloud-functions-in-go/
152 | /separate_unit_integration_tests/
141 | /working-with-db-nulls/
594 | /working-with-db-time-in-go/
```