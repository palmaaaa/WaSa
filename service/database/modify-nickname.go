package database

// This function modifies the user's nickname in the database
func (db *appdbimpl) ModifyNickname(user string, newNickname Nickname) error {
	_, err := db.c.Exec(`UPDATE users SET nickname = ? WHERE id_user = ?`, newNickname.Nickname, user)
	if err != nil {
		// Error during the execution of the query
		return err
	}
	return nil
}
