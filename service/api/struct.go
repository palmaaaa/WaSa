package api

import (
	"time"
	"wasaphoto-1849661/service/database"
)

// Error messages
const INTERNAL_ERROR_MSG = "internal server error"
const PNG_ERROR_MSG = "file is not a png format"
const JPG_ERROR_MSG = "file is not a jpg format"
const IMG_FORMAT_ERROR_MSG = "images must be jpeg or png"
const INVALID_JSON_ERROR_MSG = "invalid json format"
const INVALID_IDENTIFIER_ERROR_MSG = "identifier must be a string between 3 and 16 characters"

// JSON Error Structure
type JSONErrorMsg struct {
	Message string `json:"message"` // Error messages
}

// Photo structure for the APIs
type Photo struct {
	Comments int       `json:"comments"` // Number of comments of a photo
	Likes    int       `json:"likes"`    // Number of likes of a photo
	Owner    string    `json:"owner"`    // Unique id of the owner
	PhotoId  int       `json:"photo_id"` // Unique id of the photo
	Date     time.Time `json:"date"`     // Date in which the photo was uploaded
}

// User structure for the APIs
type User struct {
	IdUser string `json:"identifier"` // User's unique id
}

// PhotoId structure for the APIs
type PhotoId struct {
	IdPhoto int64 `json:"photo_id"` // Photo unique id
}

// Nickname structure for the APIs (useful for incoming requests)
type Nickname struct {
	Nickname string `json:"nickname"` // Nickname of a user
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

// Converts a PhotoId from the api package to a PhotoId of the database package
func (p PhotoId) ToDatabase() database.PhotoId {
	return database.PhotoId{
		IdPhoto: p.IdPhoto,
	}
}

// Converts a PhotoId from the api package to a PhotoId of the database package
func (n Nickname) ToDatabase() database.Nickname {
	return database.Nickname{
		Nickname: n.Nickname,
	}
}
