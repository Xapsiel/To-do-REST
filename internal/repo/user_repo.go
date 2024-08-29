package repo

import (
	"fmt"
	"net/http"
	"test_case/pkg/errors"
)

func (r *Repo) CreateUser(name string, password string) (int64, bool, error) {

	query := fmt.Sprintf("INSERT INTO users (login, password) VALUES ('%s','%s')", name, password)
	_, err := r.DB.Exec(query)
	if err != nil {
		fmt.Println(err)
		return -1, false, errors.New("CreateUser func", "This user already exists", http.StatusServiceUnavailable)
	}
	id, _, err := r.FindUser(name, password)
	if err != nil {
		return -1, false, err
	}
	return id, true, nil
}

func (r *Repo) FindUser(name string, password string) (int64, bool, error) {
	query := fmt.Sprintf("SElECT id  FROM users WHERE users.login='%s' AND users.password='%s'", name, password)
	rows, err := r.DB.Query(query)
	if err != nil {
		return -1, false, errors.New("FindUser func", err.Error(), http.StatusServiceUnavailable)
	}
	defer rows.Close()
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return -1, false, errors.New("FindUser func", err.Error(), http.StatusServiceUnavailable)
		}
		return id, true, nil
	}
	return -1, false, errors.New("FindUser func", "Authorization error", http.StatusNotFound)
}
func (r *Repo) FindUserById(id int) (int64, bool, error) {
	query := fmt.Sprintf("SElECT id  FROM users WHERE users.id='%d'", id)
	rows, err := r.DB.Query(query)
	if err != nil {
		return -1, false, errors.New("FindingUserById func", err.Error(), http.StatusServiceUnavailable)
	}
	defer rows.Close()
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return -1, false, errors.New("FindingUserById func", err.Error(), http.StatusServiceUnavailable)
		}
		return id, true, nil
	}
	return -1, false, errors.New("FindingUserById func", "User was not found", http.StatusNotFound)
}
