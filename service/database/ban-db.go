package database

import "errors"

// Database fuction that checks if the requesting user was banned by another 'user'. Returns 'true' if is banned, 'false' otherwise
func (db *appdbimpl) BannedUserCheck(requestingUser User, targetUser User) (bool, error) {
	var cnt int
	err := db.c.QueryRow("SELECT COUNT(*) FROM banned_users WHERE banned = ? AND banner = ?",
		requestingUser.IdUser, targetUser.IdUser).Scan(&cnt)

	if err != nil {
		return true, err
	}

	// If there's a row then the user was banned
	if cnt == 1 {
		return true, nil
	}
	return false, nil

}

// Database fuction that allows a user (banner) to ban another one (banned)
func (db *appdbimpl) BanUser(banner User, banned User) error {
	if banner == banned {
		return errors.New("users can't ban themselfes")
	}

	_, err := db.c.Exec("INSERT INTO banned_users (banner,banned) VALUES (?, ?)", banner.IdUser, banned.IdUser)
	if err != nil {
		return err
	}

	return nil
}

// Database fuction that removes a user (banned) from the banned list of another one (banner)
func (db *appdbimpl) UnbanUser(banner User, banned User) error {
	_, err := db.c.Exec("DELETE FROM banned_users WHERE (banner = ? AND banned = ?)", banner.IdUser, banned.IdUser)
	if err != nil {
		return err
	}

	return nil
}
