package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) DoLogin(username string) (int, error) {

	row := db.c.QueryRow(`SELECT id FROM User WHERE username=?;`, username)
	var id int
	exist := row.Scan(&id)
	if errors.Is(exist, sql.ErrNoRows) {
		_, err := db.c.Exec(`INSERT INTO User (username) VALUES (?);`, username)
		if err != nil {
			return id, err
		}
		row = db.c.QueryRow(`SELECT id FROM User WHERE username=?;`, username)
		err = row.Scan(&id)
		if err != nil {
			return id, nil
		}
		log.Printf("User created: %s\n", username)
	} else {
		log.Printf("User logged: %s\n", username)
	}

	return id, nil
}
