package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) GetId(username string) (int, error) {
	row := db.c.QueryRow(`SELECT id FROM User WHERE username=?`, username)
	var id int
	exist := row.Scan(&id)
	if errors.Is(exist, sql.ErrNoRows) {
		return id, ErrProfileDoesNotExist
	}
	return id, nil
}
