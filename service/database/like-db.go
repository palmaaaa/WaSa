package database

// Database function that retrieves the list of users that liked a photo
func (db *appdbimpl) GetLikesList(requestingUser User, requestedUser User, photo PhotoId) ([]User, error) {

	rows, err := db.c.Query("SELECT id_user FROM likes WHERE id_photo = ? AND id_user NOT IN (SELECT banned FROM banned_users WHERE banner = ? OR banner = ?) "+
		"AND id_user NOT IN (SELECT banner FROM banned_users WHERE banned = ?)",
		photo.IdPhoto, requestingUser.IdUser, requestedUser.IdUser, requestingUser.IdUser)
	if err != nil {
		return nil, err
	}
	// Wait for the function to finish before closing rows.
	defer func() { _ = rows.Close() }()

	// Read all the users in the resulset (users that liked the photo that didn't ban the requesting user).
	var likes []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.IdUser)
		if err != nil {
			return nil, err
		}
		likes = append(likes, user)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return likes, nil
}

/*
// Database function that gets the number of likes of a photo
func (db *appdbimpl) GetLikesLen(p PhotoId) (int, error) {

	var likes int
	err := db.c.QueryRow("SELECT COUNT(*) FROM likes WHERE (id_photo = ?)", p.IdPhoto).Scan(&likes)
	if err != nil {
		return -1, err
	}

	return likes, nil
}
*/

// Database function that adds a like of a user to a photo
func (db *appdbimpl) LikePhoto(p PhotoId, u User) error {

	_, err := db.c.Exec("INSERT INTO likes (id_photo,id_user) VALUES (?, ?)", p.IdPhoto, u.IdUser)
	if err != nil {
		return err
	}

	return nil
}

// Database function that removes a like of a user from a photo
func (db *appdbimpl) UnlikePhoto(p PhotoId, u User) error {

	_, err := db.c.Exec("DELETE FROM likes WHERE(id_photo = ? AND id_user = ?)", p.IdPhoto, u.IdUser)
	if err != nil {
		return err
	}

	return nil
}
