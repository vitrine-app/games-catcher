package main

import "C"

import (
	"github.com/Henry-Sarabia/igdb"
	"os"
)

var igdbClient = igdb.NewClient(os.Getenv("IGDB_KEY"), nil)
var db DbClient

//export GetGame
func GetGame(gameId int) *C.char {
	db = New()
	serializedGame := formatGame(getGame(gameId))
	db.Close()
	return C.CString(serializedGame)
}

//export GetFirstGame
func GetFirstGame(gameName string) *C.char {
	games, err := igdbClient.Games.Search(
		gameName,
		igdb.SetLimit(1),
	)
	if err != nil {
		panic(err.Error())
	}
	gameId := (*games[0]).ID
	return GetGame(gameId)
}

func main() {}
