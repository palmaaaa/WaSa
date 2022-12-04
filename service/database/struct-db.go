package database

import (
	"time"
)

// Photo structure for the database
type Photo struct {
	Comments int       // Number of comments of a photo
	Likes    int       // Number of likes of a photo
	Owner    string    // Unique id of the owner
	PhotoId  int       // Unique id of the photo
	Date     time.Time // Date in which the photo was uploaded
}

// User structure for the database
type User struct {
	IdUser string // User's unique id
}

// PhotoId structure for the database
type PhotoId struct {
	IdPhoto int64 `json:"photo_id"` // Photo unique id
}
