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

func formatGameFromIgdb(igdbGame igdb.Game) DbGame {
	var screenshots []string
	for _, screenshot := range igdbGame.Screenshots {
		screenshotUrl := fmt.Sprintf("https:%s", screenshot.URL)
		screenshotUrl = strings.Replace(screenshotUrl, "t_thumb", "t_1080p", -1)
		screenshots = append(screenshots, screenshotUrl)
	}
	cover := fmt.Sprintf("https:%s", igdbGame.Cover.URL)
	cover = strings.Replace(cover, "t_thumb", "t_cover_big_2x", -1)
	return DbGame{
		Name:        igdbGame.Name,
		Summary:     sql.NullString{String: igdbGame.Summary},
		Rating:      sql.NullInt64{Int64: int64(igdbGame.Rating)},
		ReleaseDate: sql.NullInt64{Int64: int64(igdbGame.FirstReleaseDate)},
		Cover:       sql.NullString{String: cover},
		Screenshots: sql.NullString{String: strings.Join(screenshots, ";")},
	}
}

func queryGame(igdbId int, apiKey string) igdb.Game {
	var client = igdb.NewClient(apiKey, nil)
	game, err := client.Games.Get(igdbId, igdb.SetFields(
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

func getGame(gameId int, apiKey string) {
	db := New()
	if db.GameExists(gameId) == false {
		game := queryGame(gameId, apiKey)
		db.AddGame(gameId, formatGameFromIgdb(game))
	}
	game := db.GetGame(gameId)
	fmt.Println(game)
	db.Close()
}
