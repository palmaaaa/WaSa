package database

// Function that inserts a like of a user, associated to a photo, in the database.
func (db *appdbimpl) LikePhoto(p PhotoId, u User) error {
	_, err := db.c.Exec("INSERT INTO likes (id_photo,id_user) VALUES (?, ?)",
		p.IdPhoto, u.IdUser)

	if err != nil {
		return err
	}

	return nil
}
