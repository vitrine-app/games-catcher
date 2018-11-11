package main

import (
	"database/sql"
	"fmt"
	"github.com/Henry-Sarabia/igdb"
	"strings"
)

type Game struct {
	Name string
}

func getGame(gameId int) {
	if db.GameExists(gameId) == false {
		game := queryIgdbGame(gameId)
		db.AddGame(formatGameFromIgdb(game, gameId))
	}
	game := db.GetGame(gameId)
	fmt.Println(game)
}

func formatGameFromIgdb(igdbGame igdb.Game, igdbId int) DbGame {
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
