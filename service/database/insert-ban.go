package database

// Fuction that bans a user (relatively to who banned him) in the db
func (db *appdbimpl) BanUser(banner User, banned User) error {
	_, err := db.c.Exec("INSERT INTO banned_users (banner,banned) VALUES (?, ?)",
		banner.IdUser, banned.IdUser)

	if err != nil {
		return err
	}

	return nil
}
