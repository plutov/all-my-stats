package stats

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	ga "google.golang.org/api/analyticsreporting/v4"
)

func getGoogleAnalyticsStats() (*ga.GetReportsResponse, error) {
	client, err := google.DefaultClient(oauth2.NoContext)
	if err != nil {
		return nil, err
	}

	svc, err := ga.New(client)
	if err != nil {
		return nil, err
	}

	req := &ga.GetReportsRequest{
		ReportRequests: []*ga.ReportRequest{
			{
				ViewId: os.Getenv("GA_ID"),
				DateRanges: []*ga.DateRange{
					{StartDate: "30daysAgo", EndDate: "today"},
				},
				Metrics: []*ga.Metric{
					{Expression: "ga:users"},
				},
			},
		},
	}

	res, err := svc.Reports.BatchGet(req).Do()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func getGoogleAnalyticsPages() (*ga.GetReportsResponse, error) {
	client, err := google.DefaultClient(oauth2.NoContext)
	if err != nil {
		return nil, err
	}

	svc, err := ga.New(client)
	if err != nil {
		return nil, err
	}

	req := &ga.GetReportsRequest{
		ReportRequests: []*ga.ReportRequest{
			{
				ViewId: os.Getenv("GA_ID"),
				DateRanges: []*ga.DateRange{
					{StartDate: "30daysAgo", EndDate: "today"},
				},
				Metrics: []*ga.Metric{
					{Expression: "ga:users"},
				},
				Dimensions: []*ga.Dimension{
					{Name: "ga:pagePath"},
				},
				OrderBys: []*ga.OrderBy{
					{FieldName: "ga:users", SortOrder: "DESCENDING"},
				},
				PageSize: 5,
			},
		},
	}

	res, err := svc.Reports.BatchGet(req).Do()
	if err != nil {
		return nil, err
	}

	return res, nil
}
