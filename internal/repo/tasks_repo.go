package repo

import (
	"fmt"
)

func Add(userID int, content string) (int, error) {
	_, _, err := FindUserById(userID)
	if err != nil {
		return -1, err
	}

	if err != nil {
		return -1, err
	}
	query := fmt.Sprintf("INSERT INTO tasks (userid,content) VALUES('%d','%s')", userID, content)
	_, err = DataBase.DB.Exec(query)
	if err != nil {
		return -1, err
	}
	id, err := findLastTasks()
	if err != nil {
		return -1, err
	}
	return int(id), nil

}
func Get(userID int) (map[int]string, error) {
	result := make(map[int]string)
	query := fmt.Sprintf("SELECT content FROM tasks WHERE userid='%d'", userID)
	rows, err := DataBase.DB.Query(query)
	if err != nil {
		return map[int]string{}, err
	}
	defer rows.Close()
	id := 1
	for rows.Next() {
		var content string
		if err := rows.Scan(&content); err != nil {
			return map[int]string{}, err
		}
		result[id] = content

		id++

	}
	return result, nil

}
func findLastTasks() (int64, error) {
	query := "SELECT max(id) from tasks"
	rows, err := DataBase.DB.Query(query)
	if err != nil {
		return -1, err
	}
	defer rows.Close()
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return -1, err
		}
	}
	return -1, err

}
