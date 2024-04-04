package database

import (
	"database/sql"
)

func (db *appdbimpl) LikePhoto(phId int, uid int) error {
	_, err := db.c.Exec(`INSERT INTO Likes (phId, uid) VALUES (?,?);`, phId, uid)
	if err != nil {
		return err
	}
	return err
}

func (db *appdbimpl) UnlikePhoto(phid int, myid int, lid int) error {
	if myid != lid {
		return ErrLike
	}
	res, err := db.c.Exec(`DELETE FROM Likes WHERE phId=? AND uid=?`, phid, myid)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		return ErrPhoto
	}
	return nil
}

func (db *appdbimpl) GetLikesPhoto(phid int) (*sql.Rows, error) {
	rows, err := db.c.Query(`SELECT u.id, u.username FROM Likes l JOIN user u ON u.id = l.uid WHERE phId=?`, phid)
	if err != nil {
		return rows, err
	}
	return rows, nil
}
