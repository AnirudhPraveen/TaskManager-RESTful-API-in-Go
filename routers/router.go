package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	// Routes for the User entity
	router = SetUserRoutes(router)
	// Routes for the User entity
	router = SetTaskRoutes(router)
	// Routes for TaskNote entity
	router = SetNoteRoutes(router)
	return router
}

