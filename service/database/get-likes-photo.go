package database

import (
	"database/sql"
)

func (db *appdbimpl) GetLikesPhoto(phid int) (*sql.Row, error) {
	res := db.c.QueryRow(`SELECT COUNT(*) FROM Likes WHERE phId=?`, phid)

	return res, nil
}
