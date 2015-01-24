package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Campaign struct {
	Campaign_id  bson.ObjectId `json:"campaign_id" bson:"_id,omitempty"`
	RootURL      string        `json:"root_url" bson:"rooturl"`
	CamapignName string        `json:"campaign_name" bson:"campaignName"`
	badges       []BadgeGroup
}
