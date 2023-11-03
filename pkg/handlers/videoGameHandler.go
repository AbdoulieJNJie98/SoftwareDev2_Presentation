package handlers

import (
	"encoding/json"
	"main/pkg/controllers"
	"main/pkg/models"
	"net/http"
)

// Strucute used to provide access to videoGameHandler methods
type VideoGameHandler struct {
	vgc controllers.VideoGameController
}

// Constructed takes in a VideoGameController varaible and assigns it to the videoGameHandler it returns
func NewVideoGameHandler(c controllers.VideoGameController) *VideoGameHandler {
	return &VideoGameHandler{
		vgc: c,
	}
}

func (vgh VideoGameHandler) PostNewVideoGame(writer http.ResponseWriter, request *http.Request) {
	var newVideoGame models.VideoGame

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(request.Body).Decode(&newVideoGame)
	if err != nil {
		http.Error(writer, "Invalid Request Body", http.StatusBadRequest)
		return
	}
	result := vgh.vgc.PostNewVideoGame(newVideoGame)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(result)

}
func (vgh VideoGameHandler) UpdateVideoGame(writer http.ResponseWriter, request *http.Request) {
	var updateVideoGame models.VideoGame

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(request.Body).Decode(&updateVideoGame)
	if err != nil {
		http.Error(writer, "Invalid Request Body", http.StatusBadRequest)
		return
	}
	result := vgh.vgc.UpdateVideoGame(updateVideoGame)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(result)

}
func (vgh VideoGameHandler) DeleteVideoGame(writer http.ResponseWriter, request *http.Request) {
	var DeleteVideoGame models.VideoGame

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(request.Body).Decode(&DeleteVideoGame)
	if err != nil {
		http.Error(writer, "Invalid Request Body", http.StatusBadRequest)
		return
	}
	result := vgh.vgc.DeleteVideoGame(DeleteVideoGame)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(result)

}
