package users

// Get a user by username
func Get() User {
	u := User{}
	// err = db.Select(&u, "SELECT * FROM users where username ='user'")
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println("users...")
	// log.Println(u)
	return u
}

// Search for users
func Search() []User {
	uu := []User{}
	// err = db.Select(&uu, "SELECT * FROM users")
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println("users...")
	// log.Println(uu)
	return uu
}
