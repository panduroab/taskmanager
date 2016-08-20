package routers

import (
	"github.com/gorilla/mux"
)

//InitRouters return mux.Router
func InitRouters() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	//Routes for the User entity
	router = SetUserRoutes(router)
	//Routes for the Task entity
	router = SetTaskRoutes(router)
	//Routes for the TaskNote entity
	router = SetNoteRoutes(router)
	return router
}
