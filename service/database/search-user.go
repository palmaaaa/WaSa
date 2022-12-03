package database

// This function filters the users by a parameter. If the parameter is missing then all registered user are returned.
// The db will look throgh any identifier or nickname (they're not necessarily the same) that match the given parameter
// (even partially) .
func (db *appdbimpl) SearchUser(nickname string) ([]User, error) {
	const query = "SELECT id_user FROM users WHERE (id_user LIKE ?) OR (nickname LIKE ?)"
	rows, err := db.c.Query(query, nickname+"%", nickname+"%")
	if err != nil {
		return nil, err
	}
	// Wait for the function to finish before closing rows.
	defer func() { _ = rows.Close() }()

	// Read all the users in the resulset.
	var res []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.IdUser)
		if err != nil {
			return nil, err
		}
		res = append(res, user)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return res, nil
}
