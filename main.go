package main

import (
	"database/sql"
	"log"
	"metroRail/Api/dbutils"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//connect to database
	db, err := sql.Open("sqlite3", "./railapi.db")

	if err != nil {
		log.Println("Driver creation failed")
	}

	//create tables
	dbutils.Initialize(db)

}
