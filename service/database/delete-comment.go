package database

// Function that removes a comment of a user, associated to a photo, in the database.
// Technically, given the structure of the db, it wouldn't be necessary to have the
// user_id to remove a comment, but it is used to make sure that whoever is requesting
// the removal is the author of the latter.
func (db *appdbimpl) UncommentPhoto(u User, c CommentId) error {
	_, err := db.c.Exec("DELETE FROM comments WHERE (id_photo = ? AND id_user= ?)",
		c.IdComment, u.IdUser)

	if err != nil {
		return err
	}

	return nil
}
