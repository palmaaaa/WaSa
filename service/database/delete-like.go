package database

// Function that removes a like of a user, associated to a photo, in the database.
func (db *appdbimpl) UnlikePhoto(p PhotoId, u User) error {
	_, err := db.c.Exec("DELETE FROM likes WHERE (id_photo = ? AND id_user = ?)",
		p.IdPhoto, u.IdUser)

	if err != nil {
		return err
	}

	return nil
}
