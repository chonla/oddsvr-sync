package sync

import (
	"fmt"

	"github.com/chonla/oddsvr-sync/database"
	"github.com/chonla/oddsvr-sync/logger"
	"github.com/chonla/oddsvr-sync/run"
	"github.com/chonla/oddsvr-sync/strava"
)

func Sync(strava *strava.Strava, db *database.Database) {
	logger.Info("Syncing data from Strava to Database")

	vr := run.NewVirtualRun(db)

	athletes := vr.AllAthleteCredentials()

	for _, athlete := range athletes {
		logger.Info(fmt.Sprintf("Syncing athlete %d with token %s", athlete.ID, athlete.AccessToken))

		lastSync := vr.GetLastSync(athlete.ID)

		logger.Debug(fmt.Sprintf("Last synchroized: %d", lastSync))

		activities, syncedOn, e := strava.GetActivities(athlete.AccessToken, lastSync) // lastSync)
		if e != nil {
			logger.Error(fmt.Sprintf("Unable to get athlete activities: %v", e))
		}

		var data []interface{}
		for _, a := range activities {
			data = append(data, a)
		}

		db.InsertBulk("activities", data)

		vr.StampLastSync(athlete.ID, syncedOn)
	}
}
