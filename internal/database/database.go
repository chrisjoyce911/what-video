package database

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

type (
	Store interface {
		Schema() error
		Migrate() error
	}

	Manage struct {
		Store   Store
		SqlInfo SqlInfo
		DB      *sqlx.DB
	}

	SqlInfo struct {
		Table   string `db:"name"`
		Version int    `db:"version"`
		Sql     string
	}
)

// Schema allows for the managment of DB tables
func (d Manage) Schema() error {
	var err error

	vt := `
	CREATE TABLE IF NOT EXISTS version (
		name VARCHAR(50) PRIMARY KEY,
		version INT NOT NULL
	  );`

	_, err = d.DB.Exec(vt)
	if err != nil {
		return err
	}

	if d.applyChange() {
		_, err = d.DB.Exec(d.SqlInfo.Sql)
		if err != nil {
			return err
		}
		log.Printf("(%s - %d) Create table applyed ...\n", d.SqlInfo.Table, d.SqlInfo.Version)
		err = d.setTableVersion()
		return err
	} else {
		log.Printf("(%s - %d) Create table skipped ...\n", d.SqlInfo.Table, d.SqlInfo.Version)
		return nil
	}
}

// Migrate changes to DB tables
func (d Manage) Migrate() error {
	var err error

	vt := `
	CREATE TABLE IF NOT EXISTS version (
		name VARCHAR(50) PRIMARY KEY,
		version INT NOT NULL
	  );`

	_, err = d.DB.Exec(vt)
	if err != nil {
		return err
	}

	if d.applyChange() {
		_, err = d.DB.Exec(d.SqlInfo.Sql)
		if err != nil {
			return err
		}
		err = d.setTableVersion()
		if err != nil {
			return err
		}
		log.Printf("(%s - %d) Table migration applyed ...\n", d.SqlInfo.Table, d.SqlInfo.Version)
		return err
	} else {
		log.Printf("(%s - %d) Table migration skipped ...\n", d.SqlInfo.Table, d.SqlInfo.Version)
		return nil
	}

}

// SetTableVersion the version of DB table
func (d Manage) setTableVersion() error {
	var err error
	sql := `
	-- Updating version record
		INSERT INTO version
			(
				name, version
			)
		VALUES
			(:name, :version)
		ON DUPLICATE KEY UPDATE
		version = :version;`

	v := struct {
		Table   string `db:"name"`
		Version int    `db:"version"`
	}{
		Table:   d.SqlInfo.Table,
		Version: d.SqlInfo.Version,
	}

	_, err = d.DB.NamedExec(sql, v)
	if err != nil {
		return err
	}
	return err
}

// Decide if the tghe update needs to applyChange
func (d Manage) applyChange() bool {
	var v int
	sql := fmt.Sprintf(`SELECT version FROM version WHERE name = "%s"`, d.SqlInfo.Table)
	err := d.DB.Get(&v, sql)
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			log.Println("sql: no rows in result set")
			v = 0
		} else {
			log.Println(err)
		}
	}
	if d.SqlInfo.Version > v {
		return true
	} else {
		return false
	}
}

// returns database NewDB connection
func NewDB(username, password, host, database string, port int) *Manage {
	//  db, err = sqlx.Connect("mysql-host", "username:password@tcp(mysql:3306)/whatvideo?charset=utf8&parseTime=True&loc=Local")
	sqlInfo := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
	)

	var dbx *sqlx.DB
	var err error
	for {
		dbx, err = sqlx.Connect("mysql", sqlInfo)

		if err != nil {
			log.Println("Connection Failed to Open")
		} else {
			log.Println("Connection Established")
			break
		}
		fmt.Println("could not connect to database, waiting 3 second")
		time.Sleep(3 * time.Second)
	}

	var m Manage
	m.DB = dbx
	return &m
}
