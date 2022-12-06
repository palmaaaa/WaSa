package database

// Database function that adds a like of a user to a photo
func (db *appdbimpl) LikePhoto(p PhotoId, u User) error {
	_, err := db.c.Exec("INSERT INTO likes (id_photo,id_user) VALUES (?, ?)",
		p.IdPhoto, u.IdUser)

	if err != nil {
		return err
	}

	return nil
}

// Database function that removes a like of a user from a photo
func (db *appdbimpl) UnlikePhoto(p PhotoId, u User) error {
	_, err := db.c.Exec("DELETE FROM likes WHERE (id_photo = ? AND id_user = ?)",
		p.IdPhoto, u.IdUser)

	if err != nil {
		return err
	}

	return nil
}
