package repo

import (
	"database/sql"
	"fmt"
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
	config.password = en.GetEnvOrDefault("DBPASSWORD", "admin")
	config.sslmode = en.GetEnvOrDefault("DBSSLMODE", "disable")
	return &Repo{DB: &sql.DB{}, configDB: configDB{name: config.name, user: config.user, password: config.password, sslmode: config.sslmode}}
}

func init() {
	envreader.Init()
	repoobj := New()
	connstr := fmt.Sprintf("user=%s password=%s sslmode=%s ", repoobj.user, repoobj.password, repoobj.sslmode)
	dbName := repoobj.name
	err := repoobj.createDB(dbName, connstr)
	if err != nil {

	}
	connstr += "dbname=" + dbName
	query := "CREATE TABLE IF NOT EXISTS users(id SERIAL PRIMARY KEY, login TEXT NOT NULL UNIQUE,password TEXT NOT NULL) "

	err = repoobj.createTable(dbName, connstr, query)
	if err != nil {

	}
	query = "CREATE TABLE IF NOT EXISTS tasks(id SERIAL PRIMARY KEY,userid INTEGER REFERENCES users (id) NOT NULL,content TEXT NOT NULL) "
	err = repoobj.createTable(dbName, connstr, query)
	if err != nil {

	}
	Repoobj = repoobj
	// Optionally, you could check if the database is reachabl
}

func (r *Repo) createDB(dbName, connstr string) error {
	var err error
	r.DB, err = sql.Open("postgres", connstr)
	if err != nil {
		return errors.New("createDB func", err.Error())

	}
	dublicate := fmt.Sprintf("pq: database \"%s\" already exists", dbName)
	_, err = r.DB.Exec("create database " + dbName)
	if err != nil && err.Error() != dublicate {
		return errors.New("createDB func", err.Error())
	}
	return nil
}

func (r *Repo) createTable(dbName, connstr, query string) error {
	var err error

	r.DB, err = sql.Open("postgres", connstr)
	if err != nil {
		return errors.New("createTable func", err.Error())
	}
	_, err = r.DB.Exec(query)
	if err != nil {
		return errors.New("createTable func", err.Error())
	}
	return nil
}
