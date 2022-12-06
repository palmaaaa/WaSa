package database

// Fuction that adds a follower to a user in the db
func (db *appdbimpl) FollowUser(follower User, followed User) error {
	_, err := db.c.Exec("INSERT INTO followers (follower,followed) VALUES (?, ?)",
		follower.IdUser, followed.IdUser)

	if err != nil {
		return err
	}

	return nil
}
