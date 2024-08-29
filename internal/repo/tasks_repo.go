package repo

import (
	"fmt"
	"net/http"
	"test_case/pkg/errors"
)

func (r *Repo) Add(userID int, content string) (int, error) {
	_, err := r.FindUserById(userID)
	if err != nil {
		return -1, err
	}

	query := fmt.Sprintf("INSERT INTO tasks (userid,content) VALUES('%d','%s')", userID, content)
	_, err = r.DB.Exec(query)
	if err != nil {
		return -1, errors.New("Add func", err.Error(), http.StatusServiceUnavailable)
	}
	id, err := r.findLastTasks()
	if err != nil {
		return -1, err
	}
	return int(id), nil

}
func (r *Repo) Get(userID int) (map[int]string, error) {
	result := make(map[int]string)
	query := fmt.Sprintf("SELECT content FROM tasks WHERE userid='%d'", userID)
	rows, err := r.DB.Query(query)
	if err != nil {
		return map[int]string{}, errors.New("Get func", err.Error(), http.StatusServiceUnavailable)
	}
	defer rows.Close()
	id := 1
	for rows.Next() {
		var content string
		if err := rows.Scan(&content); err != nil {
			return map[int]string{}, errors.New("Get func", err.Error(), http.StatusServiceUnavailable)
		}
		result[id] = content

		id++

	}
	if len(result) == 0 {
		return map[int]string{}, errors.New("Get func", "User was not found", http.StatusNotFound)
	}
	return result, nil

}
func (r *Repo) findLastTasks() (int64, error) {
	query := "SELECT max(id) from tasks"
	rows, err := r.DB.Query(query)
	if err != nil {
		return -1, errors.New("findLastTasks func()", err.Error(), http.StatusServiceUnavailable)
	}
	defer rows.Close()
	var id int64

	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return -1, errors.New("findLastTasks func()", err.Error(), http.StatusServiceUnavailable)
		}
	}
	return id, nil

}
