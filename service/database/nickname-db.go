package database

// Database function that gets a user's nickname
func (db *appdbimpl) GetNickname(user User) (string, error) {

	var nickname string

	err := db.c.QueryRow(`SELECT nickname FROM users WHERE id_user = ?`, user.IdUser).Scan(&nickname)
	if err != nil {
		// Error during the execution of the query
		return "", err
	}
	return nickname, nil
}

// Database function that modifies a user's nickname
func (db *appdbimpl) ModifyNickname(user User, newNickname Nickname) error {

	_, err := db.c.Exec(`UPDATE users SET nickname = ? WHERE id_user = ?`, newNickname.Nickname, user.IdUser)
	if err != nil {
		// Error during the execution of the query
		return err
	}
	return nil
}
