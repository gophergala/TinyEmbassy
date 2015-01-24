/*
* @Author: souravray
* @Date:   2015-01-24 12:23:16
* @Last Modified by:   souravray
* @Last Modified time: 2015-01-24 16:36:02
 */

package controllers

import (
	"log"
	"net/http"
)

func RedirectImage(w http.ResponseWriter, r *http.Request) {
	log.Println("https://s3.amazonaws.com/mazibucket/imageData/testCamapign/gopher.png")
	dispatchRedirect(w, r, "https://s3.amazonaws.com/mazibucket/imageData/testCamapign/gopher.png")
}

func RedirectTargetURL(w http.ResponseWriter, r *http.Request) {
	log.Println("http://gophergala.com/prizes"/)
	dispatchRedirect(w, r, "http://gophergala.com/prizes/")
}

func DataPoint(w http.ResponseWriter, r *http.Request) {

}
