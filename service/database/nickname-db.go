package database

// Database function that modifies a user's nickname
func (db *appdbimpl) ModifyNickname(user User, newNickname Nickname) error {

	_, err := db.c.Exec(`UPDATE users SET nickname = ? WHERE id_user = ?`, newNickname.Nickname, user.IdUser)
	if err != nil {
		// Error during the execution of the query
		return err
	}
	return nil
}
