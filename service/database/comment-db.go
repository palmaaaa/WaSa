package database

// Database function that retrieves the list of comments of a photo (minus the comments from users that banned the requesting user)
func (db *appdbimpl) GetCompleteCommentsList(requestingUser User, requestedUser User, photo PhotoId) ([]CompleteComment, error) {

	rows, err := db.c.Query("SELECT * FROM comments WHERE id_photo = ? AND id_user NOT IN (SELECT banned FROM banned_users WHERE banner = ? OR banner = ?) "+
		"AND id_user NOT IN (SELECT banner FROM banned_users WHERE banned = ?)",
		photo.IdPhoto, requestingUser.IdUser, requestedUser.IdUser, requestingUser.IdUser)
	if err != nil {
		return nil, err
	}

	// Wait for the function to finish before closing rows
	defer func() { _ = rows.Close() }()

	// Read all the comments in the resulset (comments of the photo with authors that didn't ban the requesting user).
	var comments []CompleteComment
	for rows.Next() {
		var comment CompleteComment
		err = rows.Scan(&comment.IdComment, &comment.IdPhoto, &comment.IdUser, &comment.Comment)
		if err != nil {
			return nil, err
		}

		// Get the nickname of the user that commented
		nickname, err := db.GetNickname(User{IdUser: comment.IdUser})
		if err != nil {
			return nil, err
		}
		comment.Nickname = nickname

		comments = append(comments, comment)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return comments, nil
}

/*
// Database function that gets the number of comments of a photo
func (db *appdbimpl) GetCommentsLen(p PhotoId) (int, error) {

	var comments int
	err := db.c.QueryRow("SELECT COUNT(*) FROM comments WHERE (id_photo = ?)",
		p.IdPhoto).Scan(&comments)
	if err != nil {
		return 0, err
	}

	return comments, nil
}
*/

// Database function that adds a comment of a user to a photo
func (db *appdbimpl) CommentPhoto(p PhotoId, u User, c Comment) (int64, error) {

	res, err := db.c.Exec("INSERT INTO comments (id_photo,id_user,comment) VALUES (?, ?, ?)",
		p.IdPhoto, u.IdUser, c.Comment)
	if err != nil {
		// Error executing query
		return -1, err
	}

	commentId, err := res.LastInsertId()
	if err != nil {
		// Error getting id returned by last db operation (commentId)
		return -1, err
	}

	return commentId, nil
}

/*
Technically, given the structure of the db, it wouldn't be necessary to have the
id_user to remove a comment, but it is used to make sure that whoever is requesting
the removal is the author of the latter.
Similarly for the id_photo part, except this time we want to make sure that if the url
is not valid but that comment exists for the given user, it won't be deleted.
*/

// Database function that removes a comment of a user from a photo
func (db *appdbimpl) UncommentPhoto(p PhotoId, u User, c CommentId) error {

	_, err := db.c.Exec("DELETE FROM comments WHERE (id_photo = ? AND id_user = ? AND id_comment = ?)",
		p.IdPhoto, u.IdUser, c.IdComment)
	if err != nil {
		return err
	}

	return nil
}

// Database function that removes a comment of a user from a photo ( by post author)
func (db *appdbimpl) UncommentPhotoAuthor(p PhotoId, c CommentId) error {

	_, err := db.c.Exec("DELETE FROM comments WHERE (id_photo = ? AND id_comment = ?)",
		p.IdPhoto, c.IdComment)
	if err != nil {
		return err
	}

	return nil
}
