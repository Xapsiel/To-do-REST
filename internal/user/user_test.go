package user

import (
	"reflect"
	"test_case/internal/repo"
	"testing"
)

func TestUser(t *testing.T) {
	t.Run("SignUp", TestUser_SignUp)
	t.Run("SignIn", TestUser_SignIn)
}
func TestUser_SignUp(t *testing.T) {
	repo.Repoobj.RecreateDB("test_case")
	type fields struct {
		ID       int
		Login    string
		Password string
		repo     *repo.Repo
	}
	tests := []struct {
		name    string
		fields  fields
		want    *User
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "SignUp new user", fields: fields{Login: "User", Password: "qwerty", repo: repo.New()}, want: &User{Login: "User", ID: 1, repo: nil}, wantErr: false},
		{name: "SignUp old user", fields: fields{Login: "User", Password: "qwerty", repo: repo.New()}, want: &User{Login: "", ID: 0, repo: nil}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:       tt.fields.ID,
				Login:    tt.fields.Login,
				Password: tt.fields.Password,
				repo:     tt.fields.repo,
			}
			got, err := u.SignUp()
			if (err != nil) != tt.wantErr {
				t.Errorf("User.SignUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.SignUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_SignIn(t *testing.T) {
	type fields struct {
		ID       int
		Login    string
		Password string
		repo     *repo.Repo
	}
	tests := []struct {
		name    string
		fields  fields
		want    *User
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "SignIn existing user", fields: fields{Login: "User", Password: "qwerty", repo: repo.New()}, want: &User{Login: "User", ID: 1, repo: nil}, wantErr: false},
		{name: "SignIn user with invalid password", fields: fields{Login: "User", Password: "qwerty2024", repo: repo.New()}, want: &User{}, wantErr: true},
		{name: "SignUp non-existent user", fields: fields{Login: "JohnDoe", Password: "qwerty", repo: repo.New()}, want: &User{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:       tt.fields.ID,
				Login:    tt.fields.Login,
				Password: tt.fields.Password,
				repo:     tt.fields.repo,
			}
			got, err := u.SignIn()
			if (err != nil) != tt.wantErr {
				t.Errorf("User.SignIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.SignIn() = %v, want %v", got, tt.want)
			}
		})
	}
}
