package database

// Database fuction that allows a user (banner) to ban another one (banned)
func (db *appdbimpl) BanUser(banner User, banned User) error {
	_, err := db.c.Exec("INSERT INTO banned_users (banner,banned) VALUES (?, ?)",
		banner.IdUser, banned.IdUser)

	if err != nil {
		return err
	}

	return nil
}

// Database fuction that removes a user (banned) from the banned list of another one (banner)
func (db *appdbimpl) UnbanUser(banner User, banned User) error {
	_, err := db.c.Exec("DELETE FROM banned_users WHERE(banner = ? AND banned = ?) VALUES (?, ?)",
		banner.IdUser, banned.IdUser)

	if err != nil {
		return err
	}

	return nil
}
