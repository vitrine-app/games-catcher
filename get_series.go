package main

import (
	"github.com/Henry-Sarabia/igdb"
)

func getSeries(seriesId int) DbSeries {
	if db.SeriesExists(seriesId) == false {
		series := queryIgdbSeries(seriesId)
		db.AddSeries(formatSeriesFromIgdb(series, seriesId))
	}
	return db.GetSeries(seriesId)
}

func formatSeriesFromIgdb(igdbSeries igdb.Collection, igdbId int) DbSeries {
	return DbSeries{
		IgdbId: uint(igdbId),
		Name:   igdbSeries.Name,
	}
}

func queryIgdbSeries(igdbId int) igdb.Collection {
	series, err := igdbClient.Collections.Get(igdbId, igdb.SetFields("name"))
	if err != nil {
		panic(err.Error())
	}
	return *series
}
