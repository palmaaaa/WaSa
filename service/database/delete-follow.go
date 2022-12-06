package database

// Fuction that removes a follower from a user in the db
func (db *appdbimpl) UnfollowUser(follower User, followed User) error {
	_, err := db.c.Exec("DELETE FROM followers WHERE(follower = ? AND followed = ?) VALUES (?, ?)",
		follower.IdUser, followed.IdUser)

	if err != nil {
		return err
	}

	return nil
}
