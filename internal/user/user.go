package user

import (
	"fmt"
	"test_case/internal/repo"

	"crypto/sha1"
)

type User struct {
	ID       int    `json:"id"`
	Login    string `json:"name"`
	Password string `json:"password"`
}

const (
	salt = "tklw12hfoiv3pjihu5u521jofc29urji"
)

// func hashpasswd(password string) string {
// hash := sha1.New()
// hash.Write([]byte(password))
// return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
// }

func SignUp(name string, password string) (*User, error) {
	id, err := createUser(name, hashpasswd(password))
	if err != nil {
		return &User{}, err
	}
	return &User{ID: id, Login: name}, nil
}

func createUser(name string, password string) (int, error) {
	id, ok, err := repo.CreateUser(name, password)
	if !ok {
		return -1, err
	}
	return int(id), nil
}
func SignIn(name string, password string) (*User, error) {
	id, _, err := repo.FindUser(name, hashpasswd(password))
	if err != nil {
		return &User{}, err

	}
	return &User{ID: int(id), Login: name}, nil
}
func hashpasswd(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
