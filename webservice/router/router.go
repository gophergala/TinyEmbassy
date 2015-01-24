/*
* @Author: souravray
* @Date:   2015-01-24 11:30:37
* @Last Modified by:   souravray
* @Last Modified time: 2015-01-24 13:54:49
 */

package router

import (
	"github.com/gophergala/tinyembassy/webservice/controllers"
	"github.com/gorilla/mux"
	//"net/http"
)

func Routes(rtr *mux.Router) {

	// web routes
	// apiSubrtr := rtr.PathPrefix("/web").Subrouter()
	// apiSubrtr.HandleFunc("/advertiser", controllers.SignIn).Methods("POST").Name("LogIn")")

	//api routes
	apiSubrtr := rtr.PathPrefix("/track").Subrouter()
	apiSubrtr.HandleFunc("/img/{camp}/{refr}/{badge}", controllers.RedirectImage).Methods("GET").Name("RedirectImage")
	apiSubrtr.HandleFunc("/click/{camp}/{refr}/{badge}", controllers.RedirectTargetURL).Methods("GET").Name("RedirectTargetURL")
	apiSubrtr.HandleFunc("/action/{marker}", controllers.DataPoint).Methods("GET").Name("DataPoint")
}
