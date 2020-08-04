package main

import (
	"log"
	"net/http"

	"github.com/TaskManager-RESTful-API-in-Go/common"
	"github.com/TaskManager-RESTful-API-in-Go/routers"
	"github.com/codegangsta/negroni"
)

func main() {

	// Calls startup logic
	router := routers.InitRouter()
	// Create a negroni instance
	n := negroni.Classic()
	n.UseHandler(router)
	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}
	log.Println("Listening....")
	server.ListenAndServe()
}
