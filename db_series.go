package main

import "fmt"

type DbSeries struct {
	Id     uint   `json:"id"`
	IgdbId uint   `json:"igdb_id"`
	Name   string `json:"name"`
}

func (db DbClient) GetSeries(igdbId int) DbSeries {
	var company DbSeries
	err := db.instance.QueryRow(
		"SELECT id, name FROM series WHERE igdb_id = ?", igdbId,
	).Scan(
		&company.Id,
		&company.Name,
	)
	if err != nil {
		panic(err.Error())
	}
	return company
}

func (db DbClient) AddSeries(company DbSeries) {
	query := fmt.Sprintf("INSERT INTO series (igdb_id, name, created_at) VALUES (%d, '%s', NOW())",
		company.IgdbId,
		company.Name,
	)
	insert, err := db.instance.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

func (db DbClient) SeriesExists(igdbId int) bool {
	var company DbCompany
	err := db.instance.QueryRow("SELECT name FROM series WHERE igdb_id = ?", igdbId).Scan(&company.Name)
	if err != nil {
		return false
	}
	return true
}
