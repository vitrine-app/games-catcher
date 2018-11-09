package main

import "C"

import (
	"encoding/json"
	"fmt"
	"github.com/Henry-Sarabia/igdb"
	"os"
)

//export GetGame
func GetGame(gameId int, apiKey string) *C.char {
	var client = igdb.NewClient(apiKey, nil)
	game, err := client.Games.Get(gameId, igdb.SetFields(
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
		fmt.Fprintln(os.Stderr, "Error happened: " + err.Error())
		return C.CString("error")
	}
	gameString, _ := json.Marshal(game)
	return C.CString(string(gameString))
}

func main() {
}
