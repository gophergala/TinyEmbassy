/*
* @Author: souravray
* @Date:   2015-01-24 10:34:10
* @Last Modified by:   souravray
* @Last Modified time: 2015-01-24 10:37:41
 */

package controllers

import (
	"fmt"
	"github.com/gophergala/tinyembassy/site/models"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"net/http"
)

func Login(rw http.ResponseWriter, req *http.Request) {
	websession, _ := store.Get(req, "pp-session")
	email := req.FormValue("email")
	pass := req.FormValue("password")

	encryptedPass := pass //TODO encrypt password
	s, err := mgo.Dial(conf.DbURI)

	c := s.DB(conf.DbName).C("advertiser")

	defer s.Close()

	result := models.Advertiser{}
	err = c.Find(bson.M{"email": email}).One(&result)
	if err != nil {
		fmt.Printf("user not found...", err)
		render(rw, "error.html")
	} else {
		if result.Pass == encryptedPass {
			fmt.Println("advertiser exist, password matched..")

			// Set some session values.
			websession.Values["id"] = result
			// websession.Values["test"] = "test session value"
			websession.Save(req, rw)
			fmt.Println(rw)
			render(rw, "landing.html")
		} else {
			fmt.Println("password does not match")
			render(rw, "error.html")
		}
	}
	return
}

func Logout(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Logout....")
	websession, _ := store.Get(req, "pp-session")
	websession.Options.MaxAge = -1
	websession.Save(req, rw)
}

func Signup(rw http.ResponseWriter, req *http.Request) {
	email := req.FormValue("email")
	pass := req.FormValue("password")
	name := req.FormValue("name")
	fmt.Println(email + pass + name)
	ecrypedPasswd := pass  //TODO: encrypt password
	emailVerified := false //TODO:

	s, err := mgo.Dial(conf.DbURI)
	if err != nil {
		panic(err)
	}
	c := s.DB(conf.DbName).C("advertiser")

	defer s.Close()

	result := models.Advertiser{}
	err = c.Find(bson.M{"email": email}).One(&result)
	if err != nil {
		doc := models.Advertiser{Id: bson.NewObjectId(), Email: email, Name: name, EmailVerified: emailVerified, Pass: ecrypedPasswd}
		err = c.Insert(doc)
		if err != nil {
			fmt.Printf("Can't insert document: %v\n", err)
			render(rw, "error.html")
		} else {
			// Set some session values.
			websession, _ := store.Get(req, "pp-session")
			websession.Values["id"] = doc
			websession.Save(req, rw)
			fmt.Println(websession)
			render(rw, "landing.html")
		}
	} else {
		fmt.Println("advertiser already exist..")
		fmt.Println(result)
		render(rw, "error.html")
	}

	return
}
