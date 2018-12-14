package strava

import (
	"fmt"
	"time"

	"github.com/chonla/oddsvr-sync/logger"

	"github.com/chonla/oddsvr-sync/httpclient"
	"github.com/chonla/oddsvr-sync/run"
)

const APIBASE = "https://www.strava.com/api/v3"

type Strava struct {
	clientID     string
	clientSecret string
}

func NewStrava(clientID, clientSecret string) *Strava {
	return &Strava{
		clientID:     clientID,
		clientSecret: clientSecret,
	}
}

func (s *Strava) GetActivities(token string, since int64) ([]run.Activity, int64, error) {
	page := 1
	perPage := 100
	now := time.Now()
	client := httpclient.NewClientWithToken(token)
	result := []run.Activity{}
	buffer := []run.Activity{}
	loop := true

	for loop {
		activities := []run.Activity{}
		query := fmt.Sprintf("after=%d&before=%d&page=%d&per_page=%d", since, now.Unix(), page, perPage)
		e := client.Get(fmt.Sprintf("%s/athlete/activities?%s", APIBASE, query), &activities)
		if e != nil {
			return nil, 0, e
		}
		if len(activities) == 0 {
			loop = false
		} else {
			buffer = append(buffer, activities...)
			page++
		}

		logger.Debug(fmt.Sprintf("New %d activity record(s) collected.", len(activities)))
	}

	for _, a := range buffer {
		if a.Type == "Run" {
			result = append(result, a)
		}
	}

	logger.Debug(fmt.Sprintf("New %d running record(s) extracted.", len(result)))

	return result, now.Unix(), nil
}
