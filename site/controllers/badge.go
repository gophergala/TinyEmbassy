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
	"log"
	"os"
)

func CreateBadge(w http.ResponseWriter, r *http.Request) {
	//TODO: Create badge
	fmt.Println("CreateBadge....")

	//TODO: extract data from request
	advertiserId := "advertiserID" //Dummy..
	campaignId := "campaignID"     //Dummy..
	badgeGrpID := "badgeGroupId"   //Dummy..

	//Upload image data
	f := aws.FileUpload{}

	/* dummy code */
	file, err := os.Open("tumbudu.png")
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

	err1, s3Url := f.UploadToS3(bytes, campaignId, badgeGrpID, Id)
	if err1 != nil {
		fmt.Println(err)
	}

	session, err := mgo.Dial("mongodb://te:te@flame.mongohq.com:27098/routesq")
	if err != nil {
		panic(err)
	}
	c := session.DB("routesq").C("advertiser")

	defer session.Close()

	result := models.Advertiser{}
	err = c.Find(bson.M{"_id": advertiserId}).One(&result)
	if err != nil {
		log.Fatal(err)
		//Create will fail..
		//TODO: ack the error to client
	} else {
		//	result.CreateBadge(campaignId, badgeGrpID, Id, s3Url)
		fmt.Println(campaignId, badgeGrpID, Id, s3Url)

	}

	session.Close()
	render(w, "landing.html")
	return
}

func GetBadgeData(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("GetBadgeData....")
	//TODO: return badge info for the given campaign
}
