package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Henry-Sarabia/igdb"
	"strings"
)

type FormattedGame struct {
	Name        string   `json:"name"`
	Series      string   `json:"series"`
	Genres      []string `json:"genres"`
	ReleaseDate int64    `json:"releaseDate"`
	Summary     string   `json:"summary"`
	Developer   string   `json:"developer"`
	Publisher   string   `json:"publisher"`
	Rating      int64    `json:"rating"`
	Cover       string   `json:"cover"`
	Screenshots []string `json:"screenshots"`
}

func getGame(gameId int) DbGame {
	if db.GameExists(gameId) == false {
		game := queryIgdbGame(gameId)
		insertGame(game, gameId)
	}
	game := db.GetGame(gameId)
	return game
}

func formatGame(game DbGame) string {
	var screenshots []string
	if game.Screenshots.Valid {
		screenshots = strings.Split(game.Screenshots.String, ";")
	}
	formattedGame := FormattedGame{
		Name:        game.Name,
		Genres:      db.GetGenresNameByGameId(game.Id),
		ReleaseDate: validateNumber(game.ReleaseDate),
		Summary:     validateString(game.Summary),
		Rating:      validateNumber(game.Rating),
		Cover:       validateString(game.Cover),
		Screenshots: screenshots,
	}
	if game.SeriesId.Valid {
		series := db.GetSeriesById(game.SeriesId.Int64)
		formattedGame.Series = series.Name
	}
	if game.DeveloperId.Valid {
		developer := db.GetCompanyById(game.DeveloperId.Int64)
		formattedGame.Developer = developer.Name
	}
	if game.PublisherId.Valid {
		publisher := db.GetCompanyById(game.PublisherId.Int64)
		formattedGame.Publisher = publisher.Name
	}
	gameString, err := json.Marshal(formattedGame)
	if err != nil {
		panic(err.Error())
	}
	return string(gameString)
}

func insertGame(igdbGame igdb.Game, igdbId int) DbGame {
	var screenshots []string
	for _, screenshot := range igdbGame.Screenshots {
		screenshotUrl := fmt.Sprintf("https:%s", screenshot.URL)
		screenshotUrl = strings.Replace(screenshotUrl, "t_thumb", "t_1080p", -1)
		screenshots = append(screenshots, screenshotUrl)
	}
	cover := fmt.Sprintf("https:%s", igdbGame.Cover.URL)
	cover = strings.Replace(cover, "t_thumb", "t_cover_big_2x", -1)
	dbGame := DbGame{
		IgdbId:      uint(igdbId),
		Name:        igdbGame.Name,
		Summary:     sql.NullString{String: igdbGame.Summary},
		Rating:      sql.NullInt64{Int64: int64(igdbGame.Rating)},
		ReleaseDate: sql.NullInt64{Int64: int64(igdbGame.FirstReleaseDate)},
		Cover:       sql.NullString{String: cover},
		Screenshots: sql.NullString{String: strings.Join(screenshots, ";")},
	}
	if igdbGame.Collection != 0 {
		series := getSeries(igdbGame.Collection)
		dbGame.SeriesId = sql.NullInt64{Int64: int64(series.Id)}
	}
	if len(igdbGame.Developers) != 0 {
		developer := getCompany(igdbGame.Developers[0])
		dbGame.DeveloperId = sql.NullInt64{Int64: int64(developer.Id)}
	}
	if len(igdbGame.Publishers) != 0 {
		publisher := getCompany(igdbGame.Publishers[0])
		dbGame.PublisherId = sql.NullInt64{Int64: int64(publisher.Id)}
	}
	id := db.AddGame(dbGame)
	for _, genreId := range igdbGame.Genres {
		genre := getGenre(genreId)
		db.AddGameGenre(id, genre.Id)
	}
	return dbGame
}

func queryIgdbGame(igdbId int) igdb.Game {
	game, err := igdbClient.Games.Get(igdbId, igdb.SetFields(
		"name",
		"summary",
		"collection",
		"rating",
		"developers",
		"publishers",
		"genres",
		"first_release_date",
		"cover",
		"screenshots",
	))
	if err != nil {
		panic(err.Error())
	}
	return *game
}
