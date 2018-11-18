package main

import "fmt"

type DbGenre struct {
	Id     uint64 `json:"id"`
	IgdbId uint   `json:"igdb_id"`
	Name   string `json:"name"`
}

func (db DbClient) GetGenre(igdbId int) DbGenre {
	var genre DbGenre
	err := db.instance.QueryRow(
		"SELECT id, name FROM genres WHERE igdb_id = ?", igdbId,
	).Scan(
		&genre.Id,
		&genre.Name,
	)
	if err != nil {
		panic(err.Error())
	}
	return genre
}

func (db DbClient) GetGenresNameByGameId(gameId uint) []string {
	results, err := db.instance.Query(
		"SELECT name FROM genres INNER JOIN games_genres ON genres.id = games_genres.genre_id WHERE games_genres.game_id = ?",
		gameId,
	)
	if err != nil {
		panic(err.Error())
	}
	var genres []string
	for results.Next() {
		var genre *string
		err := results.Scan(&genre)
		if err != nil {
			panic(err.Error())
		}
		genres = append(genres, *genre)
	}
	return genres
}

func (db DbClient) AddGenre(genre DbGenre) {
	query := fmt.Sprintf("INSERT INTO genres (igdb_id, name, created_at) VALUES (%d, '%s', NOW())",
		genre.IgdbId,
		genre.Name,
	)
	insert, err := db.instance.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

func (db DbClient) GenreExists(igdbId int) bool {
	var genre DbGenre
	err := db.instance.QueryRow("SELECT name FROM genres WHERE igdb_id = ?", igdbId).Scan(&genre.Name)
	if err != nil {
		return false
	}
	return true
}

func (db DbClient) AddGameGenre(gameDbId int64, genreDbId uint64) {
	query := fmt.Sprintf("INSERT INTO games_genres (game_id, genre_id) VALUES (%d, %d)",
		gameDbId,
		genreDbId,
	)
	insert, err := db.instance.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}
