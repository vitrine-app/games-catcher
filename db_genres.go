package main

import "fmt"

type DbGenre struct {
	Id     uint   `json:"id"`
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
