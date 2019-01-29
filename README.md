# all-my-stats

Google Cloud Function to display various social stats (pretty much specific for me).

- Google Analytics: visitors (1 day, 1 week, 1 month), top 5 pages (1 day, 1 week, 1 month)
- YouTube channel: delta subscribers (1 day, 1 week, 1 month), total count, top 5 videos (1 day, 1 week, 1 month)
- GitHub: total count
- Twitter: total count

## Requirements

1. Enable [Google Analytics Reporting API](https://console.developers.google.com/apis/library/analyticsreporting.googleapis.com)
2. Create YouTube API Key
3. Enable [YouTube Analytics API](https://console.developers.google.com/apis/library/youtubeanalytics.googleapis.com)
4. Create Service Account and download json key
5. Add Service Account to Google Analytics

p.s. We can get GitHub / Twitter followers count without API keys.