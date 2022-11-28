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

func (u User) ToDatabase() database.User{
	return database.User{
		IdUser: u.IdUser,
	}
}