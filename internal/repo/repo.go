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
	config := configDB{
		name:     en.GetEnvOrDefault("DBNAME", "test_case"),
		user:     en.GetEnvOrDefault("DBUSER", "postgres"),
		password: en.GetEnvOrDefault("DBPASSWORD", "ButterSQL_3301"),
		sslmode:  en.GetEnvOrDefault("DBSSLMODE", "disable"),
	}
	fmt.Println(config)

	connstr := fmt.Sprintf("user=%s password=%s sslmode=%s", config.user, config.password, config.sslmode)
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	Repoobj = &Repo{DB: db, configDB: config}
	return Repoobj
}
func init() {
	envreader.Init()
	repoobj := New()

	dbName := repoobj.name
	if err := repoobj.createDB(dbName); err != nil {
		log.Println(err)
		return
	}

	// Убедитесь, что база данных доступна
	if err := repoobj.DB.Ping(); err != nil {
		log.Fatalf("Database is not reachable: %v", err)
	}

	// Создание таблиц
	queries := []string{
		"CREATE TABLE IF NOT EXISTS users(id SERIAL PRIMARY KEY, login TEXT NOT NULL UNIQUE, password TEXT NOT NULL)",
		"CREATE TABLE IF NOT EXISTS tasks(id SERIAL PRIMARY KEY, userid INTEGER REFERENCES users (id) NOT NULL, content TEXT NOT NULL)",
	}

	for _, query := range queries {
		if err := repoobj.createTable(query); err != nil {
			log.Println(err)
			return
		}
	}
	Repoobj = repoobj
}

func (r *Repo) createDB(dbName string) error {
	dublicate := fmt.Sprintf("pq: database \"%s\" already exists", dbName)
	_, err := r.DB.Exec("CREATE DATABASE " + dbName)
	if err != nil && err.Error() != dublicate {
		return errors.New("createDB func", err.Error(), http.StatusServiceUnavailable)
	}
	return nil
}

func (r *Repo) createTable(query string) error {
	_, err := r.DB.Exec(query)
	if err != nil {
		return errors.New("createTable func", err.Error(), http.StatusServiceUnavailable)
	}
	return nil
}
