package database

// This function modifies the user's nickname on the database
func (db *appdbimpl) ModifyNickname(u User, newNickname string) error {
	_, err := db.c.Query("UPDATE users SET nickname=? WHERE id_user=?", newNickname, u.IdUser)
	if err != nil {
		// Error during the execution of the query
		return err
	}
	return nil
}
