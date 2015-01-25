/*
* @Author: souravray
* @Date:   2015-01-24 10:47:03
* @Last Modified by:   souravray
* @Last Modified time: 2015-01-25 02:35:47
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
	rtr.PathPrefix("/pub/").Handler(http.StripPrefix("/pub/", http.FileServer(http.Dir("./static/publishersBadge"))))

	//Home page
	rtr.HandleFunc("/", controllers.TLanding).Methods("GET").Name("Homepage")
	rtr.HandleFunc("/login", controllers.TLogin).Methods("GET").Name("TLogin")
	rtr.HandleFunc("/signUp", controllers.TSignUp).Methods("GET").Name("TSignUp")
	rtr.HandleFunc("/logout", controllers.TLogout).Methods("GET").Name("TLogout")

	// Advertiser routes
	apiSubrtr := rtr.PathPrefix("/advertiser").Subrouter()
	apiSubrtr.HandleFunc("/signup", controllers.Signup).Methods("POST").Name("Signup")
	apiSubrtr.HandleFunc("/login", controllers.Login).Methods("POST").Name("Login")
	apiSubrtr.HandleFunc("/logout", controllers.Logout).Methods("POST").Name("Logout")

	apiSubrtr1 := rtr.PathPrefix("/campaign").Subrouter()
	apiSubrtr1.HandleFunc("/createCampaign", controllers.CreateCampaign).Methods("POST").Name("CreateCampaign")
	apiSubrtr1.HandleFunc("/getCampaignData", controllers.GetCampaignData).Methods("POST").Name("GetCampaignData")
	apiSubrtr1.HandleFunc("/createCampaignPG", controllers.CPG).Methods("GET").Name("CPG")

	apiSubrtr.HandleFunc("/createBadge", controllers.CreateBadge).Methods("POST").Name("CreateBadge")
	apiSubrtr.HandleFunc("/getBadgeData", controllers.GetBadgeData).Methods("POST").Name("GetBadgeData")

	// web routes
	// apiSubrtr := rtr.PathPrefix("/web").Subrouter()
	// apiSubrtr.HandleFunc("/advertiser", controllers.SignIn).Methods("POST").Name("LogIn")")

	// api routes
	// apiSubrtr := rtr.PathPrefix("/api").Subrouter()
	// apiSubrtr.HandleFunc("/campaings", controllers.SignIn).Methods("GET").Name("GetCampaignList")")
}
