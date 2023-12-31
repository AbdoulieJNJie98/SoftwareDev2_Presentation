// This file is used as middle man between the handler and database that can be used to modify what is relaid back to the handler
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
func (vgc VideoGameController) UpdateVideoGame(updateVideoGame models.VideoGame) error {
	err := vgc.vgd.UpdateVideoGame(updateVideoGame)
	return err
}
func (vgc VideoGameController) DeleteVideoGame(DeleteVideoGame models.VideoGame) error {
	err := vgc.vgd.DeleteVideoGame(DeleteVideoGame)
	return err
}
