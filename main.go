package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/panduroab/taskmanager/common"
	"github.com/panduroab/taskmanager/routers"
)

//Entry point of the program
func main() {
	//calls startup logic
	common.StartUp()
	//Get the mux router object
	router := routers.InitRouters()
	//Create a negroni instance
	n := negroni.Classic()
	n.UseHandler(router)
	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}
	log.Println("Listening: ", common.AppConfig.Server)
	server.ListenAndServe()
}
