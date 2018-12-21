package run

import (
	"github.com/chonla/oddsvr-sync/database"
	"github.com/globalsign/mgo/bson"
)

type VirtualRun struct {
	db *database.Database
}

type Vr struct {
	ID              bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedBy       uint32        `json:"created_by" bson:"created_by"`
	CreatedDateTime string        `json:"created_datetime" bson:"created_datetime"`
	Title           string        `json:"title" bson:"title"`
	Detail          string        `json:"detail" bson:"detail"`
	Period          []string      `json:"period" bson:"period"`
	Link            string        `json:"link" bson:"link"`
	Engagements     []Engagement  `json:"engagements" bson:"engagements"`
}

type Engagement struct {
	AthleteID uint32  `json:"athlete_id" bson:"athlete_id"`
	Distance  float64 `json:"distance" bson:"distance"`
}

func NewVirtualRun(db *database.Database) *VirtualRun {
	return &VirtualRun{
		db: db,
	}
}

func (v *VirtualRun) GetInProgressRun() []Vr {
	return []Vr{}
}

func (v *VirtualRun) AllAthleteCredentials() []AthleteCredential {
	creds := []AthleteCredential{}
	tokens := []InvertedToken{}
	e := v.db.List("athlete", bson.M{}, &tokens)
	if e == nil {
		for _, t := range tokens {
			creds = append(creds, AthleteCredential{
				ID:           t.ID,
				AccessToken:  t.AccessToken,
				RefreshToken: t.RefreshToken,
				Expiry:       t.Expiry,
			})
		}
	}
	return creds
}

func (v *VirtualRun) StampLastSync(id uint32, stamp int64) {
	v.db.Upsert("sync", bson.M{
		"_id": id,
	}, bson.M{
		"_id":   id,
		"stamp": stamp,
	})
}

func (v *VirtualRun) GetLastSync(id uint32) int64 {
	output := map[string]interface{}{}
	e := v.db.Get("sync", bson.M{
		"_id": id,
	}, output)
	if e != nil {
		return 0
	}
	return output["stamp"].(int64)
}

func (v *VirtualRun) UpdateToken(token AthleteCredential) error {
	invToken := InvertedToken{}
	e := v.db.Get("athlete", bson.M{
		"_id": token.ID,
	}, &invToken)

	if e != nil {
		return e
	}

	invToken.Token.Expiry = token.Expiry
	invToken.Token.AccessToken = token.AccessToken
	invToken.Token.RefreshToken = token.RefreshToken

	return v.db.Replace("athlete", bson.M{
		"_id": token.ID,
	}, invToken)
}
