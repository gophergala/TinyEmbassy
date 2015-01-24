/*
* @Author: souravray
* @Date:   2015-01-24 10:47:03
* @Last Modified by:   souravray
* @Last Modified time: 2015-01-24 10:57:59
 */

package router

import (
	"github.com/gophergala/tinyembassy/site/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func Routes(rtr *mux.Router) {
	//static assets
	rtr.PathPrefix("/script/").Handler(http.StripPrefix("/script/", http.FileServer(http.Dir("./static/script"))))
	rtr.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("./static/img"))))
	rtr.PathPrefix("/style/").Handler(http.StripPrefix("/style/", http.FileServer(http.Dir("./static/style"))))
	rtr.PathPrefix("/template/").Handler(http.StripPrefix("/template/", http.FileServer(http.Dir("./static/template"))))

	//Home page
	rtr.HandleFunc("/", controllers.Landing).Methods("GET").Name("Homepage")

	// web routes
	// apiSubrtr := rtr.PathPrefix("/web").Subrouter()
	// apiSubrtr.HandleFunc("/advertiser", controllers.SignIn).Methods("POST").Name("LogIn")")

	// api routes
	// apiSubrtr := rtr.PathPrefix("/api").Subrouter()
	// apiSubrtr.HandleFunc("/campaings", controllers.SignIn).Methods("GET").Name("GetCampaignList")")
}
