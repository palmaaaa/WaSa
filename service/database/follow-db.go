package database

// Database function that adds a follower to a user
func (db *appdbimpl) FollowUser(follower User, followed User) error {
	_, err := db.c.Exec("INSERT INTO followers (follower,followed) VALUES (?, ?)",
		follower.IdUser, followed.IdUser)

	if err != nil {
		return err
	}

	return nil
}

// Database function that removes a follower from a user
func (db *appdbimpl) UnfollowUser(follower User, followed User) error {
	_, err := db.c.Exec("DELETE FROM followers WHERE(follower = ? AND followed = ?) VALUES (?, ?)",
		follower.IdUser, followed.IdUser)

	if err != nil {
		return err
	}

	return nil
}
