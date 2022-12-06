package database

// Database function that creates a photo on the database and returns the unique photo id
func (db *appdbimpl) CreatePhoto(p Photo) (int64, error) {
	res, err := db.c.Exec("INSERT INTO photos (id_user,comments,likes,date) VALUES (?,?,?,?)",
		p.Owner, p.Comments, p.Likes, p.Date)

	if err != nil {
		// Error executing query
		return -1, err
	}

	photoId, err := res.LastInsertId()
	if err != nil {
		// Error getting id returned by last db operation (photoId)
		return -1, err
	}

	return photoId, nil
}

// Database function that removes a photo from the database
func (db *appdbimpl) RemovePhoto(p PhotoId) error {
	_, err := db.c.Exec("DELETE FROM photos WHERE id_photo=?",
		p.IdPhoto)
	if err != nil {
		// Error during the execution of the query
		return err
	}

	return nil
}
