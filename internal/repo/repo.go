package repo

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	DB *sql.DB
}

var DataBase *Database = &Database{DB: &sql.DB{}}

func init() {
	connstr := "user=postgres password=password sslmode=disable "
	var err error
	DataBase.DB, err = sql.Open("postgres", connstr)
	if err != nil {
		log.Fatalf(err.Error())
	}
	dbName := "test_case"
	dublicate := fmt.Sprintf("pq: database \"%s\" already exists", dbName)
	_, err = DataBase.DB.Exec("create database " + dbName)
	if err != nil && err.Error() != dublicate {
		log.Fatalf(err.Error())
	}
	connstr += "dbname=" + dbName
	DataBase.DB, err = sql.Open("postgres", connstr)
	query := "CREATE TABLE IF NOT EXISTS users(id SERIAL PRIMARY KEY, login TEXT NOT NULL UNIQUE,password TEXT NOT NULL) "
	_, err = DataBase.DB.Exec(query)
	if err != nil {
		log.Fatalf(err.Error())
	}
	query = "CREATE TABLE IF NOT EXISTS tasks(id SERIAL PRIMARY KEY,userid INTEGER REFERENCES users (id) NOT NULL,content TEXT NOT NULL) "
	_, err = DataBase.DB.Exec(query)
	if err != nil {
		log.Fatalf(err.Error())
	}

	// Optionally, you could check if the database is reachabl
}
