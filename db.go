package main

import (
	"database/sql"
	"fmt"
	"os"
)
import _ "github.com/go-sql-driver/mysql"

type DbClient struct {
	instance *sql.DB
}

type DbGame struct {
	IgdbId      uint           `json:"igdb_id"`
	Name        string         `json:"name"`
	Summary     sql.NullString `json:"summary"`
	Rating      sql.NullInt64  `json:"rating"`
	ReleaseDate sql.NullInt64  `json:"release_date"`
	Cover       sql.NullString `json:"cover"`
	Screenshots sql.NullString `json:"screenshots"`
}

func New() DbClient {
	databaseUri := fmt.Sprintf("root:%s@tcp(%s)/vitrine", os.Getenv("MYSQL_ROOT_PASSWORD"), os.Getenv("MYSQL_HOST"))
	var db, err = sql.Open("mysql", databaseUri)
	if err != nil {
		panic(err.Error())
	}
	return DbClient{instance: db}
}

func (db DbClient) Close() {
	defer db.instance.Close()
}

func (db DbClient) GetGame(igdbId int) DbGame {
	var game DbGame
	err := db.instance.QueryRow(
		"SELECT name, summary, rating, release_date, cover, screenshots FROM games WHERE igdb_id = ?", igdbId,
	).Scan(
		&game.Name,
		&game.Summary,
		&game.Rating,
		&game.ReleaseDate,
		&game.Cover,
		&game.Screenshots,
	)
	if err != nil {
		panic(err.Error())
	}
	return game
}

func (db DbClient) AddGame(igdbId int, game DbGame) {
	query := fmt.Sprintf("INSERT INTO games (igdb_id, name, summary, rating, release_date, cover, screenshots, created_at) VALUES (%d, '%s', '%s', %d, %d, '%s', '%s', NOW())",
		igdbId,
		game.Name,
		game.Summary.String,
		game.Rating.Int64,
		game.ReleaseDate.Int64,
		game.Cover.String,
		game.Screenshots.String,
	)
	insert, err := db.instance.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

func (db DbClient) GameExists(igdbId int) bool {
	var game DbGame
	err := db.instance.QueryRow("SELECT name FROM games WHERE igdb_id = ?", igdbId).Scan(&game.Name)
	if err != nil {
		return false
	}
	return true
}
