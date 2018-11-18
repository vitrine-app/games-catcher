package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
)

type DbClient struct {
	instance *sql.DB
}

func New() DbClient {
	databaseUri := fmt.Sprintf("root:%s@tcp(%s)/vitrine", os.Getenv("MYSQL_ROOT_PASSWORD"), os.Getenv("MYSQL_HOST"))
	var db, err = sql.Open("mysql", databaseUri)
	if err != nil {
		panic(err.Error())
	}
	return DbClient{instance: db}
}

func (db DbClient) Close() {
	defer db.instance.Close()
}

func foreignKey(fieldId sql.NullInt64) string {
	if fieldId.Valid {
		return strconv.Itoa(int(fieldId.Int64))
	}
	return "NULL"
}
