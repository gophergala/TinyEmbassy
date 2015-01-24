/*
* @Author: souravray
* @Date:   2015-01-24 11:30:37
* @Last Modified by:   souravray
* @Last Modified time: 2015-01-24 12:07:38
 */

package router

import (
	// "github.com/gophergala/tinyembassy/webservice/controllers"
	"github.com/gorilla/mux"
	// "net/http"
)

func Routes(rtr *mux.Router) {

	// web routes
	// apiSubrtr := rtr.PathPrefix("/web").Subrouter()
	// apiSubrtr.HandleFunc("/advertiser", controllers.SignIn).Methods("POST").Name("LogIn")")

	// api routes
	// apiSubrtr := rtr.PathPrefix("/api").Subrouter()
	// apiSubrtr.HandleFunc("/campaings", controllers.SignIn).Methods("GET").Name("GetCampaignList")")
}
