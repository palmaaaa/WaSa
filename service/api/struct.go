package api

import (
	"wasaphoto-1849661/service/database"
)

type JSONErrorMsg struct{
	Message string `json:"message"`
}

/*
Represents the photo structure 
*/
type Photo struct{
	Comments int `json:"comments"`
	Likes int `json:"likes"`
	Owner string `json:"owner"`
	PhotoId string `json:"photo_id"`
	Date string `json:"date"`
}

type User struct{
	IdUser string `json:"identifier"`
}

// Converts a User from the api package to a User of the database package
func (u User) ToDatabase() database.User{
	return database.User{
		IdUser: u.IdUser,
	}
}

// Converts a Photo from the api package to a Photo of the database package
func (p Photo) ToDatabase() database.Photo{
	return database.Photo{
		Comments: p.Comments,
		Likes: p.Likes,
		Owner: p.Owner,
		PhotoId: p.PhotoId,
		Date: p.Date,
	}
}