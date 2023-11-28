package database

import "database/sql"

func (db *appdbimpl) GetAllUsers() ([]User, error) {
	var users []User
	rows, err := db.c.Query(`SELECT * FROM User`)
	if err != nil {
		return users, nil
	}

	for rows.Next() {
		var u User
		err = rows.Scan(&u.ID, &u.Name)
		if err != nil {
			return users, err
		}
		users = append(users, u)
	}
	return users, err
}

func (db *appdbimpl) GetAllProfiles() (*sql.Rows, error) {
	rows, err := db.c.Query("SELECT * FROM User")
	if err != nil {
		return rows, err
	}
	return rows, nil
}
