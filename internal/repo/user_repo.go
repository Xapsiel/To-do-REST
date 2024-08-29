package repo

import (
	"fmt"
	"net/http"
	"test_case/pkg/errors"
)

func (r *Repo) CreateUser(name string, password string) (int64, error) {

	query := fmt.Sprintf("INSERT INTO users (login, password) VALUES ('%s','%s')", name, password)
	_, err := r.DB.Exec(query)
	if err != nil {
		return -1, errors.New("CreateUser func", "This user already exists", http.StatusServiceUnavailable)
	}
	id, err := r.Login(name, password)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (r *Repo) Login(name string, password string) (int64, error) {
	query := fmt.Sprintf("SElECT id  FROM users WHERE users.login='%s' AND users.password='%s'", name, password)
	rows, err := r.DB.Query(query)
	if err != nil {
		return -1, errors.New("FindUser func", err.Error(), http.StatusServiceUnavailable)
	}
	defer rows.Close()
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return -1, errors.New("FindUser func", err.Error(), http.StatusServiceUnavailable)
		}
		return id, nil
	}
	return -1, errors.New("FindUser func", "Authorization error", http.StatusNotFound)
}
func (r *Repo) FindUserById(id int) (int64, error) {
	query := fmt.Sprintf("SElECT id  FROM users WHERE users.id='%d'", id)
	rows, err := r.DB.Query(query)
	if err != nil {
		return -1, errors.New("FindingUserById func", err.Error(), http.StatusServiceUnavailable)
	}
	defer rows.Close()
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return -1, errors.New("FindingUserById func", err.Error(), http.StatusServiceUnavailable)
		}
		return id, nil
	}
	return -1, errors.New("FindingUserById func", "User was not found", http.StatusNotFound)
}
