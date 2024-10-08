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
	id, err := u.repo.CreateUser(u.Login, hashpasswd(u.Password))
	if err != nil {
		return &User{}, err
	}
	return &User{ID: int(id), Login: u.Login}, nil
}

func (u *User) SignIn() (*User, error) {
	id, err := u.repo.Login(u.Login, hashpasswd(u.Password))
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
