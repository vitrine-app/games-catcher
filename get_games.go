package main

import (
	"fmt"
	"github.com/Henry-Sarabia/igdb"
)

type Game struct {
	Name string
}

func formatGame(dbGame DbGame) {

}

func queryGame(igdbId int, apiKey string) igdb.Game {
	var client = igdb.NewClient(apiKey, nil)
	game, err := client.Games.Get(igdbId, igdb.SetFields(
		"name",
		"summary",
		"collection",
		"total_rating",
		"developers",
		"publishers",
		"genres",
		"first_release_date",
		"screenshots",
		"cover",
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
		db.AddGame(gameId, game)
	}
	game := db.GetGame(gameId)
	fmt.Println(game)
	db.Close()
}
