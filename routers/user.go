package routers

import (
	"github.com/AnirudhPraveen/TaskManager-RESTful-API/controllers"
	"github.com/gorilla/mux"
)

func SetUserRoutes(router mux.Router) mux.Router {
	router.HandleFunc("/users/register", controllers.Register).Methods("POST")
	router.HandleFunc("/users/login", controllers.Login).Methods("POST")
	return router
}
