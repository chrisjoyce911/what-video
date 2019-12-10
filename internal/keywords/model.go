package keywords

import (
	"time"
	"what-video/internal/database"
)

type Keyword struct {
	ID        uint      `db:"id"`
	Owner     uint      `db:"owner"`
	Name      string    `db:"name"`
	CreatedAt time.Time ` db:"created_at" json:"created_at"`
}

// Migrate the keywords data storage to the current version
// creates storage table
// manages any schema migrations
func Migrate(myDB database.Manage) error {

	sql := `
	-- Create table keywords
	CREATE TABLE IF NOT EXISTS keywords (
		keyword_id INT PRIMARY KEY AUTO_INCREMENT,
		user_id INT NOT NULL,
		keyword VARCHAR(50) NOT NULL,
		FOREIGN KEY (user_id)
		REFERENCES users (user_id)
		ON DELETE CASCADE
	  );`

	myDB.SqlInfo.Table = "keywords"
	myDB.SqlInfo.Version = 1
	myDB.SqlInfo.Sql = sql
	err := myDB.Schema()
	if err != nil {
		return err
	}

	// myDB.SqlInfo.Sql = `ALTER TABLE keywords ADD owner INT NOT NULL;`
	// myDB.SqlInfo.Version = 2
	// err = myDB.Migrate()
	// if err != nil {
	// 	return err
	// }

	return err
}
