package main

import (
	"github.com/Henry-Sarabia/igdb"
)

func getGenre(genreId int) DbGenre {
	if db.GenreExists(genreId) == false {
		genre := queryIgdbGenre(genreId)
		db.AddGenre(formatGenreFromIgdb(genre, genreId))
	}
	return db.GetGenre(genreId)
}

func formatGenreFromIgdb(igdbGenre igdb.Genre, igdbId int) DbGenre {
	return DbGenre{
		IgdbId: uint(igdbId),
		Name:   igdbGenre.Name,
	}
}

func queryIgdbGenre(igdbId int) igdb.Genre {
	genre, err := igdbClient.Genres.Get(igdbId, igdb.SetFields("name"))
	if err != nil {
		panic(err.Error())
	}
	return *genre
}
