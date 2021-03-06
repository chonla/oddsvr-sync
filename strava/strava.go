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

func (s *Strava) RefreshToken(token string) (*run.Token, error) {
	newToken := &run.Token{}
	client := httpclient.NewClient()
	data := fmt.Sprintf("client_id=%s&client_secret=%s&grant_type=refresh_token&refresh_token=%s", s.clientID, s.clientSecret, token)
	e := client.PostForm(fmt.Sprintf("%s/oauth/token", APIBASE), data, newToken)
	if e != nil {
		return nil, e
	}

	return newToken, nil
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

		activitiesCount := len(activities)
		switch activitiesCount {
		case 0:
			logger.Debug("No more activity collected.")
		case 1:
			logger.Debug("One new activity collected.")
		default:
			logger.Debug(fmt.Sprintf("%d new activities collected.", activitiesCount))
		}

	}

	for _, a := range buffer {
		if a.Type == "Run" {
			result = append(result, a)
		}
	}

	resultCount := len(result)
	switch resultCount {
	case 0:
		logger.Debug("No running activity collected.")
	case 1:
		logger.Debug("One new running activity collected.")
	default:
		logger.Debug(fmt.Sprintf("%d new running activities collected.", resultCount))
	}

	var recentStamp int64
	if resultCount > 0 {
		recent, _ := time.Parse(time.RFC3339, result[0].StartDate)
		recentStamp = recent.Unix()
	}

	return result, recentStamp, nil
}
