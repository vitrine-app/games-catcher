package main

import "C"
import (
	"github.com/Henry-Sarabia/igdb"
	"os"
)

//export GetGame
/*func GetGame(gameId int, apiKey string) *C.char {
	game := getGame(gameId, apiKey)
	gameString, _ := json.Marshal(game)
	return C.CString(string(gameString))
}*/

var igdbClient = igdb.NewClient(os.Getenv("IGDB_KEY"), nil)
var db DbClient

func main() {
}
