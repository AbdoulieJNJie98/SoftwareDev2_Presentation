package main

import (
	"main/pkg/controllers"
	"main/pkg/database"
	"main/pkg/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	test := "10"
	println("Hello World! " + test)
	vgd := database.NewVideoGameDatabase("localhost", 5432, "user", "password", "video_game_list")
	vgc := controllers.NewVideoGameController(*vgd)
	vgh := handlers.NewVideoGameHandler(*vgc) // *vgc: the * is used to take the literal value instead of the memeory location

	// routes created by back end service for http request
	router := mux.NewRouter()
	router.HandleFunc("/videoGame", vgh.PostNewVideoGame).Methods("POST")
	router.HandleFunc("/videoGame", vgh.UpdateVideoGame).Methods("PATCH")
	router.HandleFunc("/videoGame", vgh.DeleteVideoGame).Methods("DELETE")

	// handler that designates which route the request will take, and which port to activate
	http.Handle("/", router)
	http.ListenAndServe(":8000", nil)
}
