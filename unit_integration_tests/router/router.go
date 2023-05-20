package router

import (
	"github.com/abelwhite/pigsty/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	//for use case 3 and 4
	router.HandleFunc("/api/room", controller.ViewRoom).Methods("GET")
	router.HandleFunc("/api/room", controller.CreateRoom).Methods("POST")
	router.HandleFunc("/api/room/{id}", controller.DeleteOneRoom).Methods("DELETE")
	router.HandleFunc("/api/deleteallroom", controller.DeleteAllRoom).Methods("DELETE")
	//for use case 11 and 12
	router.HandleFunc("/api/pig", controller.ViewPig).Methods("GET")
	router.HandleFunc("/api/pig", controller.CreatePig).Methods("POST")
	router.HandleFunc("/api/pig/{id}", controller.DeleteOnePig).Methods("DELETE")
	router.HandleFunc("/api/deleteallpig", controller.DeleteAllPigs).Methods("DELETE")

	//for use case 11 and 12
	router.HandleFunc("/api/pigsty", controller.ViewPigsty).Methods("GET")
	router.HandleFunc("/api/pigsty", controller.CreatePigsty).Methods("POST")
	router.HandleFunc("/api/pig/{id}", controller.DeleteOnePigsty).Methods("DELETE")
	router.HandleFunc("/api/deleteallpigsty", controller.DeleteAllPigsty).Methods("DELETE")
	return router
}
