package database

import (
	"log"
	"time"
)

func (db *appdbimpl) CommentPhoto(uid int, phid int, text string) (int, error) {
	istante := time.Now()
	date := istante.Format("2006-01-02 15:04:05")
	_, err := db.c.Exec(`INSERT INTO Comment (user, photo, string, date) VALUES (?,?,?,?)`, uid, phid, text, date)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	var cid int
	row := db.c.QueryRow(`SELECT id FROM Comment WHERE user=? AND photo=? ORDER BY date DESC`, uid, phid, text)
	row.Scan(&cid)
	log.Printf("cid: %d", cid)
	return cid, nil
}
