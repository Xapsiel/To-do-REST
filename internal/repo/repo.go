package repo

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	envreader "test_case/pkg/envReader"
	"test_case/pkg/errors"

	_ "github.com/lib/pq"
)

type configDB struct {
	name     string
	user     string
	password string
	sslmode  string
}
type Repo struct {
	DB *sql.DB
	configDB
	Err error
}

var Repoobj *Repo

func New() *Repo {
	if Repoobj != nil {
		return Repoobj
	}
	en := envreader.EnvReader{}
	config := configDB{}
	config.name = en.GetEnvOrDefault("DBNAME", "test_case")
	config.user = en.GetEnvOrDefault("DBUSER", "postgres")
	config.password = en.GetEnvOrDefault("DBPASSWORD", "postgres")
	config.sslmode = en.GetEnvOrDefault("DBSSLMODE", "disable")
	fmt.Println(config)
	return &Repo{DB: &sql.DB{}, configDB: configDB{name: config.name, user: config.user, password: config.password, sslmode: config.sslmode}}
}

func init() {
	envreader.Init()
	repoobj := New()
	connstr := fmt.Sprintf("user=%s password=%s sslmode=%s ", repoobj.user, repoobj.password, repoobj.sslmode)
	dbName := repoobj.name
	err := repoobj.createDB(dbName, connstr)
	if err != nil {
		fmt.Println(err)
		if e, ok := err.(errors.Errors); ok {
			log.Println(e.Print())
		}
		return
	}
	connstr += "dbname=" + dbName
	query := "CREATE TABLE IF NOT EXISTS users(id SERIAL PRIMARY KEY, login TEXT NOT NULL UNIQUE,password TEXT NOT NULL) "

	err = repoobj.createTable(dbName, connstr, query)
	if err != nil {
		fmt.Println(err)

		if e, ok := err.(errors.Errors); ok {
			log.Println(e.Print())
		}
		return
	}
	query = "CREATE TABLE IF NOT EXISTS tasks(id SERIAL PRIMARY KEY,userid INTEGER REFERENCES users (id) NOT NULL,content TEXT NOT NULL) "
	err = repoobj.createTable(dbName, connstr, query)
	if err != nil {
		if e, ok := err.(errors.Errors); ok {
			log.Println(e.Print())
		}
		return
	}
	Repoobj = repoobj
	// Optionally, you could check if the database is reachabl
}

func (r *Repo) createDB(dbName, connstr string) error {
	var err error
	fmt.Println(1)
	r.DB, err = sql.Open("postgres", connstr)
	if err != nil {
		fmt.Println(2, err)

		return errors.New("createDB func", err.Error(), http.StatusServiceUnavailable)

	}
	dublicate := fmt.Sprintf("pq: database \"%s\" already exists", dbName)
	_, err = r.DB.Exec("create database " + dbName)
	if err != nil && err.Error() != dublicate {
		return errors.New("createDB func", err.Error(), http.StatusServiceUnavailable)
	}
	return nil
}

func (r *Repo) createTable(dbName, connstr, query string) error {
	var err error

	r.DB, err = sql.Open("postgres", connstr)
	if err != nil {
		return errors.New("createTable func", err.Error(), http.StatusServiceUnavailable)
	}
	_, err = r.DB.Exec(query)
	if err != nil {
		return errors.New("createTable func", err.Error(), http.StatusServiceUnavailable)
	}
	return nil
}
