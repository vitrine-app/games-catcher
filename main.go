package main

import "C"

import (
	"encoding/json"
	"github.com/Henry-Sarabia/igdb"
	"os"
)

var igdbClient = igdb.NewClient(os.Getenv("IGDB_KEY"), nil)
var db DbClient

//export GetGame
func GetGame(gameId int) *C.char {
	db = New()
	game := getGame(gameId)
	db.Close()
	gameString, _ := json.Marshal(game)
	return C.CString(string(gameString))
}

func main() {}
