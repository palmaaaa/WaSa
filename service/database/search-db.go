package database

// Database function that filters the users by a parameter. Any partial match is included in the result.
// Returns a list of matching users (either by nickname or identifier)
func (db *appdbimpl) SearchUser(user User) ([]User, error) {
	const query = "SELECT id_user FROM users WHERE (id_user LIKE ?) OR (nickname LIKE ?)"
	rows, err := db.c.Query(query, user.IdUser+"%", user.IdUser+"%")
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
