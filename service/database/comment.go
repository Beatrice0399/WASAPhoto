package database

import (
	"time"
)

func (db *appdbimpl) CommentPhoto(uid int, phid int, text string) (Comment, error) {
	istante := time.Now()
	date := istante.Format("2006-01-02 15:04:05")
	var c Comment
	_, err := db.c.Exec(`INSERT INTO Comment (user, photo, string, date) VALUES (?,?,?,?)`, uid, phid, text, date)
	if err != nil {
		return c, err
	}

	row := db.c.QueryRow(`SELECT  c.id, c.user, u.username, c.string, c.date FROM Comment c
							JOIN User u ON c.user=u.id			
							WHERE c.user=? AND c.photo=? ORDER BY c.date DESC`, uid, phid, text)
	err = row.Scan(&c.ID, &c.Uid, &c.User, &c.Text, &c.Date)
	if err != nil {
		return c, err
	}

	return c, nil
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
