package sync

import (
	"fmt"
	"time"

	"github.com/chonla/oddsvr-sync/database"
	"github.com/chonla/oddsvr-sync/logger"
	"github.com/chonla/oddsvr-sync/run"
	"github.com/chonla/oddsvr-sync/strava"
)

func SyncActivities(strava *strava.Strava, db *database.Database) {
	logger.Info("Syncing data from Strava to Database")

	vr := run.NewVirtualRun(db)

	athletes := vr.AllAthleteCredentials()

	for _, athlete := range athletes {
		logger.Info(fmt.Sprintf("Syncing athlete %d with token %s", athlete.ID, athlete.AccessToken))

		lastSync := vr.GetLastSync(athlete.ID)

		logger.Debug(fmt.Sprintf("Last synchroized: %d", lastSync))

		now := time.Now().Unix()
		if athlete.Expiry <= now {
			logger.Info("Token has been expired. Refresh.")
			athlete, e := refreshToken(athlete, strava, vr)
			if e != nil {
				logger.Error("Unable to refresh token, skipped")
				continue
			}
		}

		activities, syncedOn, e := strava.GetActivities(athlete.AccessToken, lastSync) // lastSync)
		if e != nil {
			logger.Error(fmt.Sprintf("Unable to get athlete activities: %v", e))
		}

		if len(activities) > 0 {
			var data []interface{}
			for _, a := range activities {
				data = append(data, a)
			}

			db.InsertBulk("activities", data)

			vr.StampLastSync(athlete.ID, syncedOn)
		}
	}
}

func refreshToken(athlete run.AthleteCredential, strava *strava.Strava, vr *run.VirtualRun) (run.AthleteCredential, error) {
	newToken, e := strava.RefreshToken(athlete.RefreshToken)
	if e != nil {
		return e
	}

	athlete.AccessToken = newToken.AccessToken
	athlete.RefreshToken = newToken.RefreshToken
	athlete.Expiry = newToken.Expiry

	return athlete, vr.UpdateToken(athlete)
}
