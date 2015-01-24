/*
* @Author: souravray
* @Date:   2015-01-24 10:34:10
* @Last Modified by:   souravray
* @Last Modified time: 2015-01-24 10:37:41
 */

package controllers

import (
	"github.com/gorilla/sessions"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("DellStudiow-Laptop-as-Desktop"))

func Signup(w http.ResponseWriter, r *http.Request) {
	//TODO: Create entry in DB
}

func Login(w http.ResponseWriter, r *http.Request) {
	//TODO: Login: validate credentials, create session
}

func Logout(w http.ResponseWriter, r *http.Request) {
	//TODO: Logout, clean session
}

func CreateBadge(w http.ResponseWriter, r *http.Request) {
	//TODO: Create badge
}

func getCampaignData(rw http.ResponseWriter, req *http.Request) {
	//TODO: return campaign data for given Advertiser
}

func GetBadgeData(rw http.ResponseWriter, req *http.Request) {
	//TODO: return badge info for the given campaign
}
