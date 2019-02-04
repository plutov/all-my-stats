package stats

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

type sofUser struct {
	Items []struct {
		Reputation int `json:"reputation"`
	} `json:"items"`
}

func getSOFReputation() (int, error) {
	httpClient := &http.Client{
		Timeout: time.Second * 5,
	}
	resp, err := httpClient.Get(fmt.Sprintf("https://api.stackexchange.com/2.2/users/%s?order=desc&sort=reputation&site=stackoverflow", os.Getenv("SOF_ID")))
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	user := sofUser{}
	if err = json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return 0, err
	}

	if len(user.Items) == 0 {
		return 0, errors.New("sof user not found")
	}

	return user.Items[0].Reputation, nil
}
