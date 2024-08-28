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
	repo     *repo.Repo
}

const (
	salt = "tklw12hfoiv3pjihu5u521jofc29urji"
)

// func hashpasswd(password string) string {
// hash := sha1.New()
// hash.Write([]byte(password))
// return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
// }

func New(name string, password string) *User {
	return &User{Login: name, Password: password, repo: repo.New()}
}
func (u *User) SignUp() (*User, error) {
	id, err := u.createUser(u.Login, hashpasswd(u.Password))
	if err != nil {
		return &User{}, err
	}
	return &User{ID: id, Login: u.Login, repo: repo.New()}, nil
}

func (u *User) createUser(name string, password string) (int, error) {
	id, ok, err := u.repo.CreateUser(name, password)
	if !ok {
		return -1, err
	}
	return int(id), nil
}
func (u *User) SignIn() (*User, error) {
	id, _, err := u.repo.FindUser(u.Login, hashpasswd(u.Password))
	if err != nil {
		return &User{}, err

	}
	return &User{ID: int(id), Login: u.Login}, nil
}
func hashpasswd(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
