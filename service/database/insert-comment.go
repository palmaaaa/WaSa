package database

// Function that inserts a comment of a user, associated to a photo, in the database.
func (db *appdbimpl) CommentPhoto(p PhotoId, u User, c Comment) (int64, error) {
	res, err := db.c.Exec("INSERT INTO comments (id_photo,id_user,comment) VALUES (?, ?, ?)",
		p.IdPhoto, u.IdUser, c.Comment)

	if err != nil {
		// Error executing query
		return -1, err
	}

	commentId, err := res.LastInsertId()
	if err != nil {
		// Error getting id returned by last db operation (commentId)
		return -1, err
	}

	return commentId, nil
}
