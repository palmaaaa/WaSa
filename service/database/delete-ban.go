package database

// Fuction that removes a ban from a user in the db
func (db *appdbimpl) UnbanUser(banner User, banned User) error {
	_, err := db.c.Exec("DELETE FROM banned_users WHERE(banner = ? AND banned = ?) VALUES (?, ?)",
		banner.IdUser, banned.IdUser)

	if err != nil {
		return err
	}

	return nil
}
