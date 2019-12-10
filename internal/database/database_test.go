package database

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestManage_applyChange(t *testing.T) {

	type fields struct {
		Store   Store
		SqlInfo SqlInfo
		DB      *sqlx.DB
	}
	tests := []struct {
		name           string
		want           bool
		currentVersion int
		fields         fields
	}{
		{
			name:           "ok to apply change",
			want:           true,
			currentVersion: 2,
			fields: fields{
				SqlInfo: SqlInfo{
					Table:   "keywords",
					Version: 3,
				},
			},
		},
		{
			name:           "not to apply change",
			want:           false,
			currentVersion: 3,
			fields: fields{
				SqlInfo: SqlInfo{
					Table:   "keywords",
					Version: 2,
				},
			},
		},
		{
			name:           "version equal",
			want:           false,
			currentVersion: 2,
			fields: fields{
				SqlInfo: SqlInfo{
					Table:   "keywords",
					Version: 2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mock, _ := sqlmock.New()
			defer mockDB.Close()
			row := sqlmock.NewRows([]string{"version"}).AddRow(tt.currentVersion)
			mock.ExpectQuery("^SELECT version FROM version WHERE name =(.+)$").WillReturnRows(row)

			d := Manage{
				Store:   tt.fields.Store,
				SqlInfo: tt.fields.SqlInfo,
				DB:      sqlx.NewDb(mockDB, "sqlmock"),
			}

			got := d.applyChange()
			assert.Equal(t, tt.want, got)

		})
	}
}

func TestManage_setTableVersion(t *testing.T) {
	type fields struct {
		Store   Store
		SqlInfo SqlInfo
		DB      *sqlx.DB
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{
		{
			name:    "ok",
			wantErr: nil,
			fields: fields{
				SqlInfo: SqlInfo{
					Table:   "keywords",
					Version: 3,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mock, _ := sqlmock.New()
			defer mockDB.Close()
			mock.ExpectExec("INSERT INTO version").WillReturnResult(sqlmock.NewResult(1, 1))

			d := Manage{
				Store:   tt.fields.Store,
				SqlInfo: tt.fields.SqlInfo,
				DB:      sqlx.NewDb(mockDB, "sqlmock"),
			}

			got := d.setTableVersion()
			assert.Equal(t, tt.wantErr, got)

		})
	}
}
