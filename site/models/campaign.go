package models

import (
	"labix.org/v2/mgo/bson"
)

type Campaign struct {
	CampaignId     bson.ObjectId     `json:"campaign_id" bson:"_id,omitempty"`
	AdvertiserId   bson.ObjectId     `json:"_" bson:"advertiser_id,omitempty"`
	RootURL        string            `json:"root_url" bson:"rooturl"`
	CamapignName   string            `json:"campaign_name" bson:"campaignName"`
	BadgeGroupList []BadgeGroupIteam `json:"badge_group_lis" bson:"badge_group_list"`
}

type BadgeGroupIteam struct {
	Id    bson.ObjectId `json:badge_group_id" bson:"badge_group_id,omitempty"`
	Title string        `json:"title" bson:"title"`
}
