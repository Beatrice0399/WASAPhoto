package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) GetLikesPhoto(phid int) (int, error) {
	res := db.c.QueryRow(`SELECT COUNT(*) FROM Likes WHERE phId=?`, phid)
	var likes int
	exist := res.Scan(&likes)
	if errors.Is(exist, sql.ErrNoRows) {
		return 0, nil
	}
	return likes, nil
}
