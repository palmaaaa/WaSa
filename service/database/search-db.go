package database

// Database function that filters the users by a parameter. Any partial match is included in the result.
// Returns a list of matching users (either by nickname or identifier)
func (db *appdbimpl) SearchUser(searcher User, userToSearch User) ([]CompleteUser, error) {

	rows, err := db.c.Query("SELECT * FROM users WHERE ((id_user LIKE ?) OR (nickname LIKE ?)) AND id_user NOT IN (SELECT banner FROM banned_users WHERE banned = ?)",
		userToSearch.IdUser+"%", userToSearch.IdUser+"%", searcher.IdUser)
	if err != nil {
		return nil, err
	}
	// Wait for the function to finish before closing rows.
	defer func() { _ = rows.Close() }()

	// Read all the users in the resulset.
	var res []CompleteUser
	for rows.Next() {
		var user CompleteUser
		err = rows.Scan(&user.IdUser, &user.Nickname)
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
