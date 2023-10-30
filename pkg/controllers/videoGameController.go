package controllers

import (
	"main/pkg/database"
	"main/pkg/models"
)

// Strucute used to provide access to videoGameController methods
type VideoGameController struct {
	vgd database.VideoGameDatabase
}

// Constructed takes in a VideoGameDatabase varaible and assigns it to the videoGameController it returns
func NewVideoGameController(d database.VideoGameDatabase) *VideoGameController {
	return &VideoGameController{
		vgd: d,
	}
}

func (vgc VideoGameController) PostNewVideoGame(newVideoGame models.VideoGame) error {
	err := vgc.vgd.PostNewVideoGame(newVideoGame)
	return err
}
func (vgc VideoGameController) UpdateVideoGame() {
	//result := vgc.vgd.UpdateVideoGame()
	//return result
}
func (vgc VideoGameController) DeleteVideoGame() {
	//result := vgc.vgd.DeleteVideoGame()
	//return result
}
