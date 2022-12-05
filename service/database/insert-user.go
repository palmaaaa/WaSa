package database

// Function that inserts a new user in the database upon registration.
func (db *appdbimpl) CreateUser(u User) error {
	_, err := db.c.Exec("INSERT INTO users (id_user,nickname) VALUES (?, ?)",
		u.IdUser, u.IdUser)

	if err != nil {
		return err
	}

	return nil
}
