package database

// Database function that gets the stream of a user (photos of people that are followed by the latter)
func (db *appdbimpl) GetStream(user User) ([]Photo, error) {

	rows, err := db.c.Query(`SELECT * FROM photos WHERE id_user IN (SELECT followed FROM followers WHERE follower = ?) ORDER BY date DESC`,
		user.IdUser)
	if err != nil {
		return nil, err
	}
	// Wait for the function to finish before closing rows
	defer func() { _ = rows.Close() }()

	// Read all the users in the resulset
	var res []Photo
	for rows.Next() {
		var photo Photo
		err = rows.Scan(&photo.PhotoId, &photo.Owner, &photo.Date) //  &photo.Comments, &photo.Likes,
		if err != nil {
			return nil, err
		}
		res = append(res, photo)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return res, nil
}

// Database function that adds a new user in the database upon registration
func (db *appdbimpl) CreateUser(u User) error {

	_, err := db.c.Exec("INSERT INTO users (id_user,nickname) VALUES (?, ?)",
		u.IdUser, u.IdUser)

	if err != nil {
		return err
	}

	return nil
}

// [EXTRA] Database function that checks if a user exists
func (db *appdbimpl) CheckUser(targetUser User) (bool, error) {

	var cnt int
	err := db.c.QueryRow("SELECT COUNT(*) FROM users WHERE id_user = ?",
		targetUser.IdUser).Scan(&cnt)

	if err != nil {
		// Count always returns a row thanks to COUNT(*), so this situation should not happen
		return true, err
	}

	// If the counter is 1 then the user exists
	if cnt == 1 {
		return true, nil
	}
	return false, nil
}
