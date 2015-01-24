package models

import (
	"gopkg.in/mgo.v2/bson"
)

type BadgeGroup struct {
	Id          bson.ObjectId `json:"group_id" bson:"_id,omitempty"`
	Title       string        `json:"title" bson:"title"`
	Description string        `json:"description" bson:"description"`
	TargetURL   string        `json:"target_url" bson:"targeturl"`
	Badges      []Badge
}
