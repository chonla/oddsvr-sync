package run

import "github.com/globalsign/mgo/bson"

func (v *VirtualRun) AthleteActivities(id uint32, period []string) ([]Activity, error) {
	activities := []Activity{}

	e := v.db.List("activities", bson.M{
		"athlete": bson.M{
			"id": id,
		},
		"startdate": bson.M{
			"$gte": period[0],
			"$lte": period[1],
		},
	}, []string{"-startdate"}, &activities)

	if e != nil {
		return nil, e
	}

	return activities, nil
}
