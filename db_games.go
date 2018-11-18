package main

import (
	"database/sql"
	"fmt"
	"strings"
)
import _ "github.com/go-sql-driver/mysql"

type DbGame struct {
	Id          uint           `json:"id"`
	IgdbId      uint           `json:"igdb_id"`
	Name        string         `json:"name"`
	Summary     sql.NullString `json:"summary"`
	Rating      sql.NullInt64  `json:"rating"`
	ReleaseDate sql.NullInt64  `json:"release_date"`
	SeriesId    sql.NullInt64  `json:"series_id"`
	DeveloperId sql.NullInt64  `json:"developer_id"`
	PublisherId sql.NullInt64  `json:"publisher_id"`
	Cover       sql.NullString `json:"cover"`
	Screenshots sql.NullString `json:"screenshots"`
}

func (db DbClient) GetGame(igdbId int) DbGame {
	var game DbGame
	err := db.instance.QueryRow(
		"SELECT id, name, summary, rating, release_date, cover, screenshots, series_id, developer_id, publisher_id FROM games WHERE igdb_id = ?",
		igdbId,
	).Scan(
		&game.Id,
		&game.Name,
		&game.Summary,
		&game.Rating,
		&game.ReleaseDate,
		&game.Cover,
		&game.Screenshots,
		&game.SeriesId,
		&game.DeveloperId,
		&game.PublisherId,
	)
	if err != nil {
		panic(err.Error())
	}
	return game
}

func (db DbClient) AddGame(game DbGame) int64 {
	query := fmt.Sprintf("INSERT INTO games (igdb_id, name, summary, rating, release_date, series_id, developer_id, publisher_id, cover, screenshots, created_at) "+
		"VALUES (%d, \"%s\", \"%s\", %d, %d, %s, %s, %s, \"%s\", \"%s\", NOW())",
		game.IgdbId,
		game.Name,
		strings.Replace(game.Summary.String, "\"", "'", -1),
		game.Rating.Int64,
		game.ReleaseDate.Int64,
		foreignKey(game.SeriesId),
		foreignKey(game.DeveloperId),
		foreignKey(game.PublisherId),
		game.Cover.String,
		game.Screenshots.String,
	)
	fmt.Println(query)
	result, err := db.instance.Exec(query)
	if err != nil {
		panic(err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	return id
}

func (db DbClient) GameExists(igdbId int) bool {
	var game DbGame
	err := db.instance.QueryRow("SELECT name FROM games WHERE igdb_id = ?", igdbId).Scan(&game.Name)
	if err != nil {
		return false
	}
	return true
}
