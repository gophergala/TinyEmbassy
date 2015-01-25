/*
* @Author: souravray
* @Date:   2015-01-24 12:23:16
* @Last Modified by:   souravray
* @Last Modified time: 2015-01-25 09:33:57
 */

package controllers

import (
	"fmt"
	"github.com/gophergala/tinyembassy/webservice/models"
	"github.com/gophergala/tinyembassy/webservice/stacker/worker"
	"github.com/gorilla/mux"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"net/http"
	"time"
)

func RedirectImage(w http.ResponseWriter, r *http.Request) {
	t := []string{"i", "u"}
	vars := mux.Vars(r)
	jobqueue.AddJob(worker.Payload{time.Now(), t, vars["badge"], vars["refr"], conf.DbURI, conf.DbName, counter.Id})
	l := getTargetUrl(vars["badge"])
	fmt.Println(l)
	dispatchRedirect(w, r, "https://s3.amazonaws.com/mazibucket/imageData/testCamapign/gopher.png")
}

func RedirectTargetURL(w http.ResponseWriter, r *http.Request) {
	t := []string{"c"}
	vars := mux.Vars(r)
	jobqueue.AddJob(worker.Payload{time.Now(), t, vars["badge"], vars["refr"], conf.DbURI, conf.DbName, counter.Id})
	dispatchRedirect(w, r, "http://gophergala.com/prizes/")
}

func DataPoint(w http.ResponseWriter, r *http.Request) {

}

func getTargetUrl(groupId string) (targetUrl string) {
	// If it is in the cache return it instantly
	if data, err := lruCatche.Get(groupId); err == nil {
		if targetUrl, ok := data.(string); ok {
			return targetUrl
		}
	}

	//incase of cache miss
	s, err := mgo.Dial(conf.DbURI)
	if err != nil {
		panic(err)
	}
	defer s.Close()
	s.SetSafe(&mgo.Safe{})
	c := s.DB(conf.DbName).C("counter")
	var group models.BadgeGroup
	err = c.FindId(bson.ObjectIdHex(groupId)).One(&group)
	if err != nil {
		targetUrl = group.TargetURL
		lruCatche.Set(groupId, targetUrl)
	}
	return targetUrl
}
