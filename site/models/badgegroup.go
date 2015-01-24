package models

import (
	"labix.org/v2/mgo/bson"
)

type BadgeGroup struct {
	BadgeGroupId bson.ObjectId `json:"group_id" bson:"_id,omitempty"`
	CampaignId   bson.ObjectId `json:"campaign_id" bson:"_id,omitempty"`
	Title        string        `json:"title" bson:"title"`
	TargetURL    string        `json:"target_url" bson:"targeturl"`
	Badges       []Badge
}
