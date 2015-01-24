package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Badge struct {
	Id       bson.ObjectId `json:badge_id" bson:"_id,omitempty"`
	IamgeURL string        `json:"imageURL, bson:"imageURL"`
	Size     ImageSize
}

type ImageSize struct {
	Height int `json:"height", bson:"height"`
	Width  int `json:"width", bosn:"width"`
}
