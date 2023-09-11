package router

import (
	"github.com/AmanMittal2804/mongodb/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	// router.HandleFunc("/api/movie/{id}",controller.).Methods("DELETE")
	// router.HandleFunc("/api/deletemovie",controller.).Methods("POST")

	return router
}
