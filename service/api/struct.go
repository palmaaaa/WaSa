package api

import (
	"wasaphoto-1849661/service/database"
)

// constants
const INTERNAL_ERROR_MSG = "Internal server error"

// JSON Error Structure
type JSONErrorMsg struct {
	Message string `json:"message"` // Error messages
}

// Photo structure for the APIs
type Photo struct {
	Comments int    `json:"comments"` // Number of comments of a photo
	Likes    int    `json:"likes"`    // Number of likes of a photo
	Owner    string `json:"owner"`    // Unique id of the owner
	PhotoId  int    `json:"photo_id"` // Unique id of the photo
	Date     string `json:"date"`     // Date in which the photo was uploaded
}

// User structure for the APIs
type User struct {
	IdUser string `json:"identifier"` // User's unique id
}

// Converts a User from the api package to a User of the database package
func (u User) ToDatabase() database.User {
	return database.User{
		IdUser: u.IdUser,
	}
}

// Converts a Photo from the api package to a Photo of the database package
func (p Photo) ToDatabase() database.Photo {
	return database.Photo{
		Comments: p.Comments,
		Likes:    p.Likes,
		Owner:    p.Owner,
		PhotoId:  p.PhotoId,
		Date:     p.Date,
	}
}
