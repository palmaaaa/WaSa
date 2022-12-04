package database

// Function creates a photo on the database and returns the photo id
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
