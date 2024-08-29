package repo

import (
	"database/sql"
	"reflect"
	"testing"
)

func TestTask_Repo(t *testing.T) {
	t.Run("CreateUser", TestRepo_CreateUser)
	t.Run("Add", TestRepo_Add)
	t.Run("Get", TestRepo_Get)
}
func TestRepo_Add(t *testing.T) {

	type fields struct {
		DB       *sql.DB
		configDB configDB
		Err      error
	}
	type args struct {
		userID  int
		content string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Add default task", fields: fields{DB: Repoobj.DB, configDB: Repoobj.configDB}, args: args{userID: 1, content: "Buy dog ​​food"}, want: 1, wantErr: false},
		{name: "Add task to non-existent user", fields: fields{DB: Repoobj.DB, configDB: Repoobj.configDB}, args: args{userID: 100, content: "Buy fish ​​food"}, want: -1, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repo{
				DB:       tt.fields.DB,
				configDB: tt.fields.configDB,
				Err:      tt.fields.Err,
			}
			got, err := r.Add(tt.args.userID, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repo.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Repo.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepo_Get(t *testing.T) {

	type fields struct {
		DB       *sql.DB
		configDB configDB
		Err      error
	}
	type args struct {
		userID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[int]string
		wantErr bool
	}{
		{name: "Get task by existing user", fields: fields{DB: Repoobj.DB, configDB: Repoobj.configDB}, args: args{userID: 1}, want: map[int]string{1: "Buy dog ​​food"}, wantErr: false},
		{name: "Get task by non-existent user", fields: fields{DB: Repoobj.DB, configDB: Repoobj.configDB}, args: args{userID: 100}, want: map[int]string{}, wantErr: true},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repo{
				DB:       tt.fields.DB,
				configDB: tt.fields.configDB,
				Err:      tt.fields.Err,
			}
			got, err := r.Get(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repo.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repo.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
