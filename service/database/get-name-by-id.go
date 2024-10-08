package database

import (
	"database/sql"
	"errors"
)

// Database function that returns the username of the given id
func (db *appdbimpl) GetNameById(id int) (string, error) {
	row := db.c.QueryRow(`SELECT username FROM User WHERE id=?`, id)
	var name string
	exist := row.Scan(&name)
	if errors.Is(exist, sql.ErrNoRows) {
		return name, ErrProfileDoesNotExist
	}
	return name, nil
}
