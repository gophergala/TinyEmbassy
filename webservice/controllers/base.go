/*
* @Author: souravray
* @Date:   2015-01-24 11:26:29
* @Last Modified by:   souravray
* @Last Modified time: 2015-01-24 19:06:54
 */

package controllers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/sessions"
	// "html/template"
	"fmt"
	"github.com/gophergala/tinyembassy/webservice/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"os"
)

type Config struct {
	DbURI  string `json:"db_uri"`
	DbName string `json:"db_name"`
}

var conf Config

var store = sessions.NewCookieStore([]byte("tim-tim-tok"))

// var templates = template.Must(template.ParseFiles(
// // "static/template/landing.html",
// ))

// func render(w http.ResponseWriter, tmpl string) {
// 	err := templates.ExecuteTemplate(w, tmpl, nil)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }

var counter models.CampaignCounter

func dispatchRedirect(w http.ResponseWriter, r *http.Request, url string) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
	w.Header().Set("Pragma", "no-cache")                                   // HTTP 1.0.
	w.Header().Set("Expires", "0")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect) // Proxies
}

func dispatchNotFound(w http.ResponseWriter, message string) {
	err := errors.New(message)
	http.Error(w, err.Error(), 404)
}

func dispatchError(w http.ResponseWriter, message string) {
	err := errors.New(message)
	http.Error(w, err.Error(), 500)
}

func dispatchJSON(w http.ResponseWriter, response interface{}) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

func init() {
	configpath := os.Getenv("TE_CONF_PATH")
	if configpath == "" {
		configpath = "/tmp/config.json"
	}
	// read condig file
	content, err := ioutil.ReadFile(configpath)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(content, &conf)
	if err != nil {
		panic(err)
	}
	initCollectiohn()
}

func initCollectiohn() {
	s, err := mgo.Dial(conf.DbURI)
	if err != nil {
		panic(err)
	}
	defer s.Close()
	s.SetSafe(&mgo.Safe{})
	c := s.DB(conf.DbName).C("counter")
	counter = models.CampaignCounter{Id: bson.NewObjectId()}

	err = c.Insert(counter)
	if err != nil {
		fmt.Printf("Can't insert document: %v\n", err)
		os.Exit(1)
	}
}
