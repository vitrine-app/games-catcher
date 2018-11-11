package main

import (
	"github.com/Henry-Sarabia/igdb"
)

func getCompany(companyId int) DbCompany {
	if db.CompanyExists(companyId) == false {
		company := queryIgdbCompany(companyId)
		db.AddCompany(formatCompanyFromIgdb(company, companyId))
	}
	return db.GetCompany(companyId)
}

func formatCompanyFromIgdb(igdbCompany igdb.Company, igdbId int) DbCompany {
	return DbCompany{
		IgdbId: uint(igdbId),
		Name:   igdbCompany.Name,
	}
}

func queryIgdbCompany(igdbId int) igdb.Company {
	game, err := igdbClient.Companies.Get(igdbId, igdb.SetFields("name"))
	if err != nil {
		panic(err.Error())
	}
	return *game
}
