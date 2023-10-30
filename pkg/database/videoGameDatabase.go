package database

import (
	"context"
	"fmt"
	"main/pkg/models"

	"github.com/jackc/pgx/v5"
)

// Struct representing a new PostgreSQL object
type VideoGameDatabase struct {
	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDBName   string
}

// Creates a new instance of the PostgreSQL struct
func NewVideoGameDatabase(host string, port int, user string, pass string, db string) *VideoGameDatabase {

	return &VideoGameDatabase{
		PostgresHost:     host,
		PostgresPort:     port,
		PostgresUser:     user,
		PostgresPassword: pass,
		PostgresDBName:   db,
	}
}

func (vgd VideoGameDatabase) PostNewVideoGame(newVideoGame models.VideoGame) error {
	// construct postgres connection string by replacing %s and %d with corrasponding varaibles
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", vgd.PostgresHost, vgd.PostgresPort, vgd.PostgresUser, vgd.PostgresPassword, vgd.PostgresDBName)

	connection, err := pgx.Connect(context.Background(), psqlconn)
	if err != nil {
		print("error:" + err.Error())
		return err
	}
	query := "insert into video_games (title, platform, genre, price) VALUES (@title, @platform, @genre, @price)"
	args := pgx.NamedArgs{
		"title":    newVideoGame.Title,
		"platform": newVideoGame.Platform,
		"genre":    newVideoGame.Genre,
		"price":    newVideoGame.Price,
	}
	_, err = connection.Exec(context.Background(), query, args) // look into context (providing an underscore in golang means it will ignore the argument)
	if err != nil {
		print("error 2:" + err.Error())
		return err
	}
	return err
}
func (vgd VideoGameDatabase) UpdateVideoGame() {
	// psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", vgd.PostgresHost, vgd.PostgresPort, vgd.PostgresUser, vgd.PostgresPassword, vgd.PostgresDBName)

	// connection, err := pgx.Connect(context.Background(), psqlconn)
	// if err != nil {
	// 	return err
	// }
	// return "Update Video Game"
}
func (vgd VideoGameDatabase) DeleteVideoGame() {
	// psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", vgd.PostgresHost, vgd.PostgresPort, vgd.PostgresUser, vgd.PostgresPassword, vgd.PostgresDBName)

	// connection, err := pgx.Connect(context.Background(), psqlconn)
	// if err != nil {
	// 	return err
	// }
	// return "Delete Video Game"
}
