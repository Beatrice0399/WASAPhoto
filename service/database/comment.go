package database

import (
	"time"
)

func (db *appdbimpl) CommentPhoto(uid int, phid int, text string) (int, error) {
	istante := time.Now()
	date := istante.Format("2006-01-02 15:04:05")
	_, err := db.c.Exec(`INSERT INTO Comment (user, photo, string, date, visible) VALUES (?,?,?,?)`, uid, phid, text, date)
	if err != nil {
		return -1, err
	}
	var cid int
	row := db.c.QueryRow(`SELECT id FROM Comment WHERE user=? AND photo=? ORDER BY date DESC`, uid, phid, text)
	row.Scan(&cid)
	return cid, nil
}

func (db *appdbimpl) UncommentPhoto(cid int, phid int, uid int) error {
	res, err := db.c.Exec(`DELETE FROM Comment WHERE id=? AND user=? AND photo=?`, cid, uid, phid)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		return err
	}
	return nil
}
