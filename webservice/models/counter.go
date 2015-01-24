/*
* @Author: souravray
* @Date:   2015-01-24 17:08:56
* @Last Modified by:   souravray
* @Last Modified time: 2015-01-24 20:52:18
 */

package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Counter struct {
	U int `json:"unique", bson:"u"`
	I int `json:"impression", bosn:"i"`
	C int `json:"click", bosn:"c"`
}

type CampaignCounter struct {
	Id bson.ObjectId `json:"id",bson:"_id,omitempty"`
	P  int
	Ct Counter
	T  map[string]TimedCounter `json:"time", bson:"time"`
}

type TimedCounter struct {
	H map[string]Counter `json:"time", bson:"hour"`
	D map[string]Counter `json:"time", bson:"day"`
	W map[string]Counter `json:"time", bson:"week"`
}
