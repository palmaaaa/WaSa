package api

//"fmt"

// function that verifies if the identifier of a user has the right lenght
func valid_identifier(identifier string) bool {
	return len(identifier) >= 3 && len(identifier) <= 16
}

// function that creates a photo directory for a certain user
func create_photo_folder(user_id string, current_dir string) (string, error) {
	return "bruh", nil
}
