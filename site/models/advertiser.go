/*
* @Author: souravray
* @Date:   2015-01-24 10:39:38
* @Last Modified by:   souravray
* @Last Modified time: 2015-01-24 10:40:56
 */

package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Advertiser struct {
	Id            bson.ObjectId `json:"-" bson:"_id,omitempty"`
	Email         string        `json:"email" bson:"email"`
	Img           string        `json:"img"  bson:"img,omitempty"`
	Name          string        `json:"name" bson:"name"`
	EmailVerified bool
	Pass          string
	Campaigns     []Campaign
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
