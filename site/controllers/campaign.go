/*
* @Author: souravray
* @Date:   2015-01-24 10:34:31
* @Last Modified by:   souravray
* @Last Modified time: 2015-01-24 10:37:43
 */

package controllers

import (
	"fmt"
	"github.com/gophergala/tinyembassy/site/models"
	// "github.com/gorilla/sessions"
	// "github.com/nu7hatch/gouuid"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"net/http"
)

func CreateCampaign(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("in Create campaign....")
	rootURL := req.FormValue("rootURL")
	camapignName := req.FormValue("camapignName")

	websession, _ := store.Get(req, "pp-session")
	fmt.Println(websession)

	advertiser := websession.Values["id"].(*models.Advertiser)
	fmt.Println(websession)
	session, err := mgo.Dial(conf.DbURI)
	if err != nil {
		panic(err)
	}
	c := session.DB(conf.DbName).C("campaigns")

	defer session.Close()

	result := models.Campaign{}
	err = c.Find(bson.M{"CamapignName": camapignName}).One(&result)
	if err != nil {
		doc := models.Campaign{CampaignId: bson.NewObjectId(), AdvertiserId: advertiser.Id, RootURL: rootURL, CamapignName: camapignName}
		err = c.Insert(doc)
		if err != nil {
			fmt.Printf("Can't insert document: %v\n", err)
			render(rw, "error.html")
		} else {
			render(rw, "landing.html")
		}
	} else {
		fmt.Println("advertiser already exist..")
		fmt.Println(result)
		render(rw, "error.html")
	}
	session.Close()
	return
}

func CPG(rw http.ResponseWriter, req *http.Request) {
	render(rw, "CPG.html")
	return
}

func CBG(rw http.ResponseWriter, req *http.Request) {
	render(rw, "CBG.html")
	return
}
func CreateBadgeGroup(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("in Create Badge Group....")
	targetURL := req.FormValue("targetURL")
	title := req.FormValue("title")
	campaignName := req.FormValue("campaignName")

	websession, _ := store.Get(req, "pp-session")
	fmt.Println(websession)

	advertiser := websession.Values["id"].(*models.Advertiser)
	fmt.Println(websession)
	session, err := mgo.Dial(conf.DbURI)
	if err != nil {
		panic(err)
	}
	c := session.DB(conf.DbName).C("campaigns")

	defer session.Close()

	campaign := models.Campaign{}
	fmt.Println("Search for " + campaignName)
	err = c.Find(bson.M{"campaignName": campaignName, "advertiser_id": advertiser.Id}).One(&campaign)
	if err == nil {
		bc := session.DB(conf.DbName).C("badgeGroup")
		badgeGroup := models.BadgeGroup{}
		err = bc.Find(bson.M{"title": title, "_campaign_id": campaign.CampaignId}).One(&badgeGroup)
		if err != nil {
			doc := models.BadgeGroup{BadgeGroupId: bson.NewObjectId(), CampaignId: campaign.CampaignId, Title: title, TargetURL: targetURL}
			err = bc.Insert(doc)
			if err != nil {
				fmt.Printf("Can't insert document: %v\n", err)
				render(rw, "error.html")
			} else {
				render(rw, "landing.html")
			}
		} else {
			fmt.Println("badge already exist...")
			render(rw, "error.html")
		}
	} else {
		fmt.Println("Campaign does not exit")
		fmt.Println(err)
		render(rw, "error.html")
	}
	session.Close()
	return
}

func GetCampaignData(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("GetCampaignData....")
	//TODO: return campaign data for given Advertiser
}
