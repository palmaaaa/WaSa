package database

import "errors"

// Database function that retrieves the list of photos of a user (only if the requesting user is not banned by that user)
func (db *appdbimpl) GetPhotosList(requestinUser User, targetUser User) ([]Photo, error) {
	banned, err := db.BannedUserCheck(requestinUser, targetUser)
	if err != nil {
		return nil, err
	}
	if banned {
		return nil, errors.New("can't get stream, the requesting user is blocked")
	}

	rows, err := db.c.Query("SELECT * FROM photos WHERE id_user = ?", targetUser.IdUser)
	if err != nil {
		return nil, err
	}
	// Wait for the function to finish before closing rows
	defer func() { _ = rows.Close() }()

	// Read all the photos in the resulset
	var photos []Photo
	for rows.Next() {
		var photo Photo
		err = rows.Scan(&photo.PhotoId, &photo.Owner, &photo.Comments, &photo.Likes, &photo.Date)
		if err != nil {
			return nil, err
		}
		photos = append(photos, photo)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return photos, nil
}

// Database function that checks if a photo exists
func (db *appdbimpl) CheckPhotoExistence(targetPhoto PhotoId) bool {
	var rows int
	err := db.c.QueryRow("SELECT COUNT(*) FROM photos WHERE (id_photo = ?)", targetPhoto.IdPhoto).Scan(&rows)
	if err != nil {
		return false
	}

	if rows == 0 {
		return false
	}
	return true

}

// Database function that retrieves a specific photo (only if the requesting user is not banned by that owner of that photo).
// If the requesting user is blocked by the
func (db *appdbimpl) GetPhoto(requestinUser User, targetPhoto PhotoId) (Photo, error) {

	if !db.CheckPhotoExistence(targetPhoto) {
		return Photo{}, ErrPhotoDoesntExist
	}

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
