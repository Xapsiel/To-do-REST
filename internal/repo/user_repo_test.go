package repo

import (
	"database/sql"
	"testing"
)

func TestUser_Repo(t *testing.T) {
	t.Run("CreateUser", TestRepo_CreateUser)
	t.Run("Login", TestRepo_Login)
}
func TestRepo_CreateUser(t *testing.T) {
	Repoobj.RecreateDB("test_case")

	type fields struct {
		DB       *sql.DB
		configDB configDB
		Err      error
	}
	type args struct {
		name     string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "SignUp new user", fields: fields{DB: Repoobj.DB, configDB: Repoobj.configDB}, args: args{name: "User", password: "qwerty"}, want: 1, wantErr: false},
		{name: "SignUp old user", fields: fields{DB: Repoobj.DB, configDB: Repoobj.configDB}, args: args{name: "User", password: "qwerty"}, want: -1, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repo{
				DB:       tt.fields.DB,
				configDB: tt.fields.configDB,
				Err:      tt.fields.Err,
			}
			got, err := r.CreateUser(tt.args.name, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repo.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Repo.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepo_Login(t *testing.T) {

	type fields struct {
		DB       *sql.DB
		configDB configDB
		Err      error
	}
	type args struct {
		name     string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Find existing user", fields: fields{DB: Repoobj.DB, configDB: Repoobj.configDB}, args: args{name: "User", password: "qwerty"}, want: 1, wantErr: false},
		{name: "Find user with invalid password", fields: fields{DB: Repoobj.DB, configDB: Repoobj.configDB}, args: args{name: "User", password: "qwerty2024"}, want: -1, wantErr: true},
		{name: "Find non-existent user", fields: fields{DB: Repoobj.DB, configDB: Repoobj.configDB}, args: args{name: "JohnDoe", password: "qwerty"}, want: -1, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repo{
				DB:       tt.fields.DB,
				configDB: tt.fields.configDB,
				Err:      tt.fields.Err,
			}
			got, err := r.Login(tt.args.name, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repo.FindUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Repo.FindUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
