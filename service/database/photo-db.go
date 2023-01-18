package database

// Database function that retrieves the list of photos of a user (only if the requesting user is not banned by that user)
func (db *appdbimpl) GetPhotosList(requestingUser User, targetUser User) ([]Photo, error) { // requestinUser User,

	rows, err := db.c.Query("SELECT * FROM photos WHERE id_user = ? ORDER BY date DESC", targetUser.IdUser)
	if err != nil {
		return nil, err
	}
	// Wait for the function to finish before closing rows
	defer func() { _ = rows.Close() }()

	// Read all the photos in the resulset
	var photos []Photo
	for rows.Next() {
		var photo Photo
		err = rows.Scan(&photo.PhotoId, &photo.Owner, &photo.Date)
		if err != nil {
			return nil, err
		}

		comments, err := db.GetCompleteCommentsList(requestingUser, targetUser, PhotoId{IdPhoto: int64(photo.PhotoId)}) // Old: GetCommentsLen
		if err != nil {
			return nil, err
		}
		photo.Comments = comments

		likes, err := db.GetLikesList(requestingUser, targetUser, PhotoId{IdPhoto: int64(photo.PhotoId)}) // Old: GetLikesLen
		if err != nil {
			return nil, err
		}
		photo.Likes = likes

		photos = append(photos, photo)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return photos, nil
}

// Database function that retrieves a specific photo (only if the requesting user is not banned by that owner of that photo).
func (db *appdbimpl) GetPhoto(requestinUser User, targetPhoto PhotoId) (Photo, error) {

	var photo Photo
	err := db.c.QueryRow("SELECT * FROM photos WHERE (id_photo = ?) AND id_user NOT IN (SELECT banner FROM banned_user WHERE banned = ?)",
		targetPhoto.IdPhoto, requestinUser.IdUser).Scan(&photo)

	if err != nil {
		return Photo{}, ErrUserBanned
	}

	return photo, nil

}

// Database function that creates a photo on the database and returns the unique photo id
func (db *appdbimpl) CreatePhoto(p Photo) (int64, error) {

	res, err := db.c.Exec("INSERT INTO photos (id_user,date) VALUES (?,?)",
		p.Owner, p.Date)

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

/*
Adding the owner is an additional security measure to delete photos that are actually owned
by that user
*/

// Database function that removes a photo from the database
func (db *appdbimpl) RemovePhoto(owner User, p PhotoId) error {

	_, err := db.c.Exec("DELETE FROM photos WHERE id_user = ? AND id_photo = ? ",
		owner.IdUser, p.IdPhoto)
	if err != nil {
		// Error during the execution of the query
		return err
	}

	return nil
}

// [Util] Database function that checks if a photo exists
func (db *appdbimpl) CheckPhotoExistence(targetPhoto PhotoId) (bool, error) {

	var rows int
	err := db.c.QueryRow("SELECT COUNT(*) FROM photos WHERE (id_photo = ?)", targetPhoto.IdPhoto).Scan(&rows)
	if err != nil {
		return false, err
	}

	if rows == 0 {
		return false, nil
	}
	return true, nil

}
