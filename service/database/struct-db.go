package database

import "time"

/*
// Photo structure for the database
type Photo struct {
	Comments int       // Number of comments of a photo
	Likes    int       // Number of likes of a photo
	Owner    string    // Unique id of the owner
	PhotoId  int       // Unique id of the photo
	Date     time.Time // Date in which the photo was uploaded
}
*/

// Photo structure for the database
type Photo struct {
	Comments []CompleteComment `json:"comments"` // Array of comments of the photo
	Likes    []CompleteUser    `json:"likes"`    // Array of useres that liked the photo
	Owner    string            `json:"owner"`    // Unique id of the owner
	PhotoId  int               `json:"photo_id"` // Unique id of the photo
	Date     time.Time         `json:"date"`     // Date in which the photo was uploaded
}

// User structure for the database
type User struct {
	IdUser string `json:"user_id"` // User's unique id
}

// User structure for the database
type CompleteUser struct {
	IdUser   string `json:"user_id"`  // User's unique id
	Nickname string `json:"nickname"` // Nickname of a user
}

// PhotoId structure for the database
type PhotoId struct {
	IdPhoto int64 `json:"photo_id"` // Photo unique id
}

// Nickname structure for the database
type Nickname struct {
	Nickname string `json:"nickname"` // Nickname of a user
}

// Comment structure for the database
type Comment struct {
	Comment string `json:"comment"` // Comment content
}

// CommentId structure for the database
type CommentId struct {
	IdComment int64 `json:"comment_id"` // Identifier of a comment
}

// CompleteComment structure for the database
type CompleteComment struct {
	IdComment int64  `json:"comment_id"` // Identifier of a comment
	IdPhoto   int64  `json:"photo_id"`   // Photo unique id
	IdUser    string `json:"user_id"`    // User's unique id
	Nickname  string `json:"nickname"`   // Nickname of a user
	Comment   string `json:"comment"`    // Comment content
}
