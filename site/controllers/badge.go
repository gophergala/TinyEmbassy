/*
* @Author: souravray
* @Date:   2015-01-24 10:35:13
* @Last Modified by:   souravray
* @Last Modified time: 2015-01-24 10:37:45
 */

package controllers

import (
	"fmt"
	"github.com/gophergala/tinyembassy/site/models"
	// "github.com/gorilla/sessions"
	"github.com/nu7hatch/gouuid"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"net/http"

	"bufio"
	"github.com/gophergala/tinyembassy/site/controllers/aws"
	"io"
	// "io/ioutil"
	// "log"
	"os"
)

func CreateBadge(rw http.ResponseWriter, req *http.Request) {
	//TODO: Create badge
	fmt.Println("CreateBadge....")
	websession, _ := store.Get(req, "pp-session")

	// extract data from request
	advertiser := websession.Values["id"].(*models.Advertiser)
	campaigntitle := req.FormValue("campaigntitle")
	badgeGroupName := req.FormValue("badgeGroupName")

	//Upload image data
	f := aws.FileUpload{}

	file1, _, err := req.FormFile("uploadedfile")

	if err != nil {
		fmt.Fprintln(rw, err)
		return
	}

	defer file1.Close()

	out, err := os.Create("/tmp/uploadedfile")
	if err != nil {
		fmt.Fprintf(rw, "Unable to create the file for writing. Check your write access privilege")
		return
	}
	// write the content from POST to the file
	_, err = io.Copy(out, file1)
	if err != nil {
		fmt.Fprintln(rw, err)
	}

	out.Close()

	file, err := os.Open("/tmp/uploadedfile")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	//get file size
	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	bytes := make([]byte, size)

	//Prepare buffer to post
	buffer := bufio.NewReader(file)
	_, err = buffer.Read(bytes)
	/* dummy code till here... */

	// func (this *FileUpload) UploadToS3(data []byte, campaign string, badge string) (err error) {
	//Generate unique id for uploaded file
	var u5 *uuid.UUID
	u5, err = uuid.NewV4()
	if err != nil {
		// return  err
		fmt.Println(err)
	}
	Id := u5.String()

	// result := models.Advertiser{}
	// err = c.Find(bson.M{"_id": advertiserId}).One(&result)
	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	//	result.CreateBadge(campaignId, badgeGrpID, Id, s3Url)
	// 	fmt.Println(campaignId, badgeGrpID, Id, s3Url)

	// }
	session, err := mgo.Dial(conf.DbURI)
	c := session.DB(conf.DbName).C("campaigns")

	campaign := models.Campaign{}
	fmt.Println("Search for " + campaigntitle)
	err = c.Find(bson.M{"campaignName": campaigntitle, "advertiser_id": advertiser.Id}).One(&campaign)
	if err == nil {
		bc := session.DB(conf.DbName).C("badgeGroup")
		badgeGroup := models.BadgeGroup{}
		err = bc.Find(bson.M{"title": badgeGroupName, "_campaign_id": campaign.CampaignId}).One(&badgeGroup)
		if err == nil {

			err1, s3Url := f.UploadToS3(bytes, campaigntitle, Id)
			os.Remove("/tmp/uploadedfile")
			if err1 != nil {
				fmt.Println(err)
			}
			fmt.Println("S3 url" + s3Url)

			session, err := mgo.Dial(conf.DbURI)
			if err != nil {
				panic(err)
			}
			defer session.Close()

			badge := models.Badge{IamgeURL: s3Url, S3BadgeId: string(Id), Id: Id}
			//badgeGroup.Badges = append(badgeGroup.Badges, badge)
			_, err = bc.Upsert(bson.M{"BadgeGroupId": badgeGroup.BadgeGroupId}, bson.M{"$addToSet": bson.M{"badges": badge}})

			// b := session.DB(conf.DbName).C("badges")
			// err = bc.Find(bson.M{"title": campaigntitle, "_campaign_id": campaign.CampaignId}).One(&badgeGroup)
			// doc := models.BadgeGroup{BadgeGroupId: bson.NewObjectId(), CampaignId: campaign.CampaignId, Title: title, TargetURL: targetURL}
			// err = bc.Insert(doc)
			if err != nil {
				fmt.Printf("Can't insert document: %v\n", err)
				render(rw, "error.html")
			} else {
				render(rw, "landing.html")
			}
		} else {
			fmt.Println("badge already exist...")
			fmt.Println(err)
			render(rw, "error.html")
		}
	} else {
		fmt.Println("Campaign does not exit")
		fmt.Println(err)
		render(rw, "error.html")
	}

	session.Close()
	render(rw, "landing.html")
	return
}

func CreateBadgeT(rw http.ResponseWriter, req *http.Request) {
	render(rw, "badge.html")
	return
}

func GetBadgeData(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("GetBadgeData....")
	//TODO: return badge info for the given campaign
}
