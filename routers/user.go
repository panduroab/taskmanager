package routers

import (
	"github.com/gorilla/mux"
	"github.com/panduroab/taskmanager/controllers"
)

//SetUserRoutes set the User resource routers
func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/users/register", controllers.Register).Methods("POST")
	router.HandleFunc("/users/login", controllers.Login).Methods("POST")
	return router
}
