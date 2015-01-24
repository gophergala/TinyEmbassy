/*
* @Author: souravray
* @Date:   2015-01-24 12:23:16
* @Last Modified by:   souravray
* @Last Modified time: 2015-01-24 20:27:37
 */

package controllers

import (
	"fmt"
	//"github.com/gophergala/tinyembassy/webservice/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"os"
	"time"
)

func RedirectImage(w http.ResponseWriter, r *http.Request) {
	s, err := mgo.Dial(conf.DbURI)
	if err != nil {
		panic(err)
	}
	defer s.Close()
	s.SetSafe(&mgo.Safe{})
	c := s.DB(conf.DbName).C("counter")
	fmt.Println(counter.Id)
	t := time.Now()
	hour := fmt.Sprintf("t.h.%d%02d%02d%02d%02d%02d.i",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	err = c.Update(bson.M{"id": counter.Id}, bson.M{"$inc": bson.M{"ct.i": 1, hour: 1}})
	if err != nil {
		fmt.Printf("Can't update document %v\n", err)
		os.Exit(1)
	}

	log.Println("https://s3.amazonaws.com/mazibucket/imageData/testCamapign/gopher.png")
	dispatchRedirect(w, r, "https://s3.amazonaws.com/mazibucket/imageData/testCamapign/gopher.png")
}

func RedirectTargetURL(w http.ResponseWriter, r *http.Request) {
	log.Println("http://gophergala.com/prizes")
	dispatchRedirect(w, r, "http://gophergala.com/prizes/")
}

func DataPoint(w http.ResponseWriter, r *http.Request) {

}
