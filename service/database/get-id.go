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
		return 0, ErrProfileDoesNotExist
	}
	return id, nil
}
