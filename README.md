# all-my-stats

Google Cloud Function to display various social stats (pretty much specific for me).

- Google Analytics: users and top 5 pages (last 30 days)
- YouTube channel: subscribers count
- GitHub: followers count
- Twitter: followers count

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

[https://europe-west1-func-230310.cloudfunctions.net/stats](https://europe-west1-func-230310.cloudfunctions.net/stats)

```
```