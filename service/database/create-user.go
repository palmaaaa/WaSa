package database

import(
)

func (db *appdbimpl) CreateUser(u User) (error) {
	res, err := db.c.Exec("INSERT INTO users (id_user,nickname) VALUES (?, ?)",
	u.IdUser,u.IdUser)
	
	if err != nil {
		return err
	}

	_, err = res.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}