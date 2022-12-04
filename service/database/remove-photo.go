package database

// Function removes a photo from the database. Returns an error if anything went wrong.
func (db *appdbimpl) RemovePhoto(p PhotoId) error {
	_, err := db.c.Exec("DELETE FROM photos WHERE id_photo=?",
		p.IdPhoto)
	if err != nil {
		// Error during the execution of the query
		return err
	}

	/*
		_, err = res.RowsAffected()
		if err != nil {
			// Error getting numbers of rows affected
			return err
		}
	*/

	return nil
}
