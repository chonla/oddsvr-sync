package main

import (
	"fmt"
	"os"

	"github.com/chonla/oddsvr-sync/database"
	"github.com/chonla/oddsvr-sync/logger"
	"github.com/chonla/oddsvr-sync/strava"
	"github.com/chonla/oddsvr-sync/sync"
)

func main() {
	stravaClientID, e := env("ODDSVR_STRAVA_CLIENT_ID", "", "strava client id", true)
	if e != nil {
		logger.Error(e.Error())
		os.Exit(1)
	}
	stravaClientSecret, e := env("ODDSVR_STRAVA_CLIENT_SECRET", "", "strava client secret", true)
	if e != nil {
		logger.Error(e.Error())
		os.Exit(1)
	}
	dbServer, _ := env("ODDSVR_DB", "127.0.0.1:27017", "database address", false)

	st := strava.NewStrava(stravaClientID, stravaClientSecret)
	db, e := database.NewDatabase(dbServer, "vr")
	if e != nil {
		logger.Error(e.Error())
		os.Exit(1)
	}

	sync.SyncActivities(st, db)
	sync.SyncVirtualRuns(db)
}

func env(key, defaultValue, name string, errorIfMissing bool) (string, error) {
	value, found := os.LookupEnv(key)
	if !found || value == "" {
		if errorIfMissing {
			return "", fmt.Errorf("seems like %s (%s) is missing from environment variables", name, key)
		}
		logger.Info(fmt.Sprintf("%s is set to default value", name))
		logger.Info(fmt.Sprintf("you can override this using %s environment variable", key))
		value = defaultValue
	}
	return value, nil
}
