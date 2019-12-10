package users

import (
	"time"
	"what-video/internal/database"
)

type User struct {
	ID        uint      `db:"id"`
	UserName  string    `db:"username"`
	Email     uint      `db:"email"`
	Password  string    `db:"password_hash"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	LastLogin time.Time `db:"last_login"`
}

// Migrate the keywords data storage to the current version
// creates storage table
// manages any schema migrations
func Migrate(myDB database.Manage) error {

	sql := `
		-- Create table users
		CREATE TABLE users(
			id SERIAL PRIMARY KEY,
			username VARCHAR(50) UNIQUE NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			password_hash TEXT NOT NULL
		);`

	myDB.SqlInfo.Table = "users"
	myDB.SqlInfo.Version = 1
	myDB.SqlInfo.Sql = sql
	err := myDB.Schema()
	if err != nil {
		return err
	}

	return err
}
