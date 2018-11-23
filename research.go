package main

import (
	"encoding/json"
	"fmt"
	"github.com/Henry-Sarabia/igdb"
	"strings"
)

type ResearchedGame struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Cover string `json:"cover"`
}

func researchGames(gameName string, listSize int) string {
	games, err := igdbClient.Games.Search(
		strings.Replace(gameName, "²", "2", -1),
		igdb.SetFields("name", "cover"),
		igdb.SetLimit(listSize),
	)
	if err != nil {
		panic(err.Error())
	}
	var research []ResearchedGame
	for _, game := range games {
		cover := "https://images.igdb.com/igdb/image/upload/t_cover_big_2x/nocover_qhhlj6.jpg"
		if len(game.Cover.URL) > 0 {
			cover = strings.Replace(fmt.Sprintf("https:%s", game.Cover.URL), "t_thumb", "t_cover_big_2x", -1)
		}
		research = append(research, ResearchedGame{
			Id:    game.ID,
			Name:  game.Name,
			Cover: cover,
		})
	}
	researchString, err := json.Marshal(research)
	if err != nil {
		panic(err.Error())
	}
	return string(researchString)
}
