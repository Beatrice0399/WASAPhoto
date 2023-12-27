package database

import (
	"database/sql"
	"time"
)

func (db *appdbimpl) UploadPhoto(id int, img []byte) (Photo, error) {
	var photo Photo
	photo.Image = img
	istante := time.Now()
	date := istante.Format("2006-01-02 15:04:05")

	res, err := db.c.Exec(`INSERT INTO Photo (user, image, date) VALUES (?, ?, ?)`,
		id, img, date)
	if err != nil {
		return photo, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return photo, err
	}

	photo.ID = int(lastInsertID)
	return photo, nil
}

func (db *appdbimpl) DeletePhoto(pid int, myid int) error {
	res, err := db.c.Exec(`DELETE FROM Photo WHERE id=? AND user=?`, pid, myid)
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

func (db *appdbimpl) GetPhotoComments(phId int) (*sql.Rows, error) {
	//var comments []Comment
	rows, err := db.c.Query(`SELECT c.id, u.username, c.string, c.date FROM Comment c
								JOIN User u ON c.user=u.id WHERE photo=? AND visible=1`, phId)
	if err != nil {
		return nil, err
	}

	/*
		for rows.Next() {
			var c Comment
			err = rows.Scan(&c.ID, &c.User, &c.Text, &c.Date)
			if err != nil {
				return nil, err
			}
			//log.Printf("phid: %d, user: %s, txt: %s, date: %s\n", c.ID, c.User, c.Text, c.Date)
			comments = append(comments, c)
		}
	*/

	return rows, err
}
