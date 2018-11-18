package main

import "C"

import (
	"fmt"
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

func main() {}
