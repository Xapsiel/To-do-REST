package repo

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/lib/pq"
)

func TestRepo_createTable(t *testing.T) {
	Repoobj.RecreateDB("test_case")
	type fields struct {
		DB       *sql.DB
		configDB configDB
		Err      error
	}
	type args struct {
		dbName  string
		connstr string
		query   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "create new db", fields: fields{DB: New().DB, configDB: Repoobj.configDB}, args: args{dbName: "test_case", connstr: fmt.Sprintf("host=%s port=%s user=%s dbname=postgres password=%s sslmode=%s ", Repoobj.host, Repoobj.port, Repoobj.user, Repoobj.password, Repoobj.sslmode)}, wantErr: false},
		{name: "create existing db", fields: fields{DB: New().DB, configDB: Repoobj.configDB}, args: args{dbName: "test_case", connstr: fmt.Sprintf("host=%s port=%s user=%s dbname=postgres password=%s sslmode=%s ", Repoobj.host, Repoobj.port, Repoobj.user, Repoobj.password, Repoobj.sslmode)}, wantErr: false},
	}
	// TODO: Add test cases.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repo{
				DB:       tt.fields.DB,
				configDB: tt.fields.configDB,
				Err:      tt.fields.Err,
			}
			if err := r.createTable(tt.args.dbName, tt.args.connstr, tt.args.query); (err != nil) != tt.wantErr {
				t.Errorf("Repo.createTable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
