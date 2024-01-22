package routers

import (
	"adv/controller"
	"net/http"

	"github.com/gorilla/mux"
)


func Router() *mux.Router {
	routes := mux.NewRouter()
	//Static Files
	routes.Handle("/static/",controller.StaticFilesEmbed())
	//FORM
	routes.HandleFunc("/login",controller.Form).Methods("POST")
	//Cookies
	routes.HandleFunc("/set-cookie",controller.SetCookie).Methods("GET")
	routes.HandleFunc("/get-cookie",controller.GetCookie).Methods("GET")
	//File Server
	return routes
}

func RouteFile() *http.ServeMux {
	routes := http.NewServeMux()
	routes.Handle("/static/",http.StripPrefix("/static",controller.StaticFilesEmbed()))
	routes.HandleFunc("/",controller.TemplateFile)
	routes.HandleFunc("/home",controller.ServeFileEmbed)
	routes.HandleFunc("/about",controller.ServerFile)
	return routes
}