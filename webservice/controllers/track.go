/*
* @Author: souravray
* @Date:   2015-01-24 12:23:16
* @Last Modified by:   souravray
* @Last Modified time: 2015-01-25 06:52:35
 */

package controllers

import (

	//"github.com/gophergala/tinyembassy/webservice/models"
	"github.com/gophergala/tinyembassy/webservice/stacker/worker"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func RedirectImage(w http.ResponseWriter, r *http.Request) {
	t := []string{"i", "u"}
	vars := mux.Vars(r)
	jobqueue.AddJob(worker.Payload{time.Now(), t, vars["badge"], vars["refr"], conf.DbURI, conf.DbName, counter.Id})
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
