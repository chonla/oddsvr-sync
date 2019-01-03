package run

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

func (v *VirtualRun) AthleteActivities(id uint32, period []string) ([]Activity, error) {
	activities := []Activity{}

	from, e := time.Parse(time.RFC3339, period[0])
	if e != nil {
		return nil, e
	}
	fromLocation := from.Location()
	newForm := time.Date(from.Year(), from.Month(), from.Day(), 0, 0, 0, 0, fromLocation)
	fromStamp := newForm.Format(time.RFC3339)
	to, e := time.Parse(time.RFC3339, period[1])
	if e != nil {
		return nil, e
	}
	toLocation := to.Location()
	newTo := time.Date(to.Year(), to.Month(), to.Day(), 23, 59, 59, 999, toLocation)
	toStamp := newTo.Format(time.RFC3339)

	e = v.db.List("activities", bson.M{
		"athlete": bson.M{
			"id": id,
		},
		"startdate": bson.M{
			"$gte": fromStamp,
			"$lte": toStamp,
		},
	}, []string{"-startdate"}, &activities)

	if e != nil {
		return nil, e
	}

	return activities, nil
}
