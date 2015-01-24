/*
* @Author: souravray
* @Date:   2015-01-24 10:39:38
* @Last Modified by:   souravray
* @Last Modified time: 2015-01-24 10:40:56
 */

package models

import (
	"labix.org/v2/mgo/bson"
)

type Advertiser struct {
	Id            bson.ObjectId `json:"-" bson:"_id,omitempty"`
	Email         string        `json:"email" bson:"email"`
	Img           string        `json:"img"  bson:"img,omitempty"`
	Name          string        `json:"name" bson:"name"`
	EmailVerified bool          `json:"verified" bson:"verified"`
	Pass          string
	CampaignList  []CampaignIteams `json:"camp_list" bson:"camp_list"`
}

type CampaignIteams struct {
	Id   bson.ObjectId `json:camp_id" bson:"camp_id,omitempty"`
	Name string        `json:name" bson:"name,omitempty"`
}

func (user *Advertiser) Validator() error {
	if len(user.Email) == 0 {
		return ErrNotFilled
	}
	if !EmailRegexp.MatchString(user.Email) {
		return ErrInvalidEmail
	}
	return nil
}
