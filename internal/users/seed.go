package users

import (
	"log"
	"what-video/internal/database"
)

// Seed the keywords data storage to the current version
// creates storage table
// manages any schema migrations
func Seed(myDB database.Manage) error {

	var u User
	sql := `INSERT INTO users
		(username, email, password_hash)
		VALUES
		(:username, :email, :password_hash);`

	u.UserName = "tomm"
	u.Email = "tomm@example.net"
	hash, _ := HashPassword("mygreatpassword")
	u.Password = hash
	res, err := myDB.DB.NamedExec(sql, &u)
	if err != nil {
		return err
	}

	id, _ := res.LastInsertId()
	affected, _ := res.RowsAffected()
	log.Printf("Seed userID :%d , affected :%d", id, affected)

	return err
}
