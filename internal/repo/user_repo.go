package repo

import (
	"fmt"
)

func CreateUser(name string, password string) (int64, bool, error) {

	query := fmt.Sprintf("INSERT INTO users (login, password) VALUES ('%s','%s')", name, password)
	_, err := DataBase.DB.Exec(query)
	if err != nil {
		return -1, false, err
	}
	id, _, _ := FindUser(name, password)
	return id, true, nil
}

func FindUser(name string, password string) (int64, bool, error) {
	query := fmt.Sprintf("SElECT id  FROM users WHERE users.login='%s' AND users.password='%s'", name, password)
	rows, err := DataBase.DB.Query(query)
	if err != nil {
		return -1, false, err
	}
	defer rows.Close()
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return -1, false, err
		}
		return id, true, nil
	}
	return -1, false, fmt.Errorf("Такого пользователя нет")
}
func FindUserById(id int) (int64, bool, error) {
	query := fmt.Sprintf("SElECT id  FROM users WHERE users.id='%d'", id)
	rows, err := DataBase.DB.Query(query)
	if err != nil {
		return -1, false, err
	}
	defer rows.Close()
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return -1, false, err
		}
		return id, true, nil
	}
	return -1, false, fmt.Errorf("Такого пользователя нет")
}
