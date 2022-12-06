package database

// Function that removes a comment of a user, associated to a photo, in the database.
// Technically, given the structure of the db, it wouldn't be necessary to have the
// id_user to remove a comment, but it is used to make sure that whoever is requesting
// the removal is the author of the latter.
// Similarly for the id_photo part, except this time we want to make sure that if the url
// is not valid but that comment exists for the given user, it won't be deleted.
func (db *appdbimpl) UncommentPhoto(p PhotoId, u User, c CommentId) error {
	_, err := db.c.Exec("DELETE FROM comments WHERE (id_photo = ? AND id_user = ? AND id_comment = ?)",
		p.IdPhoto, u.IdUser, c.IdComment)

	if err != nil {
		return err
	}

	return nil
}
