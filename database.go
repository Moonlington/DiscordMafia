package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func loadDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func addCity(name string) (sql.Result, error) {
	return db.Exec("INSERT INTO Cities(Name) VALUES (?)", name)
}

func addCompany(name string, wealth, diff float64, city int) (sql.Result, error) {
	return db.Exec("INSERT INTO Companies(Name, Wealth, Difficulty, City) VALUES (?,?,?,?)", name, wealth, diff, city)
}
