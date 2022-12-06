package database

// This function gets the stream of a user.
func (db *appdbimpl) GetStream(user User) ([]Photo, error) {
	const query = `SELECT * FROM photos WHERE id_user IN (SELECT followed FROM followers WHERE followed = ?) ORDER BY date ASC`
	rows, err := db.c.Query(query, user.IdUser)
	if err != nil {
		return nil, err
	}
	// Wait for the function to finish before closing rows.
	defer func() { _ = rows.Close() }()

	// Read all the users in the resulset.
	var res []Photo
	for rows.Next() {
		var photo Photo
		err = rows.Scan(&photo)
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
