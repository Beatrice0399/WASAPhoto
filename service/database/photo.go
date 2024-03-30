package database

import (
	"database/sql"
	"time"
)

func (db *appdbimpl) NewPhoto(id int) (int, error) {
	var photo Photo
	istante := time.Now()
	date := istante.Format("2006-01-02 15:04:05")

	res, err := db.c.Exec(`INSERT INTO Photo (user, date) VALUES (?, ?)`,
		id, date)
	if err != nil {
		return -1, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	photo.ID = int(lastInsertID)
	return int(lastInsertID), nil
}

func (db *appdbimpl) UpdatePathPhoto(phid int, path string) error {
	_, err := db.c.Exec(`UPDATE Photo SET image_path = ? WHERE id = ? `, path, phid)
	if err != nil {
		return err
	}
	return nil
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

func (db *appdbimpl) GetPhoto(phId int) (Photo, error) {
	var photo Photo
	row := db.c.QueryRow(`SELECT p.id, p.user, p.image_path, p.date
							FROM Photo p 
							WHERE id = ?`, phId)

	err := row.Scan(&photo.ID, &photo.User, &photo.Path, &photo.Date)
	if err != nil {
		return photo, err
	}
	res, err := db.GetLikesPhoto(photo.ID)
	for res.Next() {
		var u User
		err = res.Scan(&u.Uid, &u.Username)
		if err != nil {
			return photo, err
		}
		photo.Likes = append(photo.Likes, u)
	}

	com, err := db.GetPhotoComments(photo.ID)
	if err != nil {
		return photo, err
	}
	defer com.Close()
	for com.Next() {
		var c Comment
		err = com.Scan(&c.ID, &c.User, &c.Text, &c.Date)
		if err != nil {
			return photo, err
		}
		photo.Comments = append(photo.Comments, c)
	}
	return photo, nil
}

func (db *appdbimpl) GetPhotoComments(phId int) (*sql.Rows, error) {
	// var comments []Comment
	rows, err := db.c.Query(`SELECT c.id, u.username, c.string, c.date FROM Comment c
								JOIN User u ON c.user=u.id WHERE photo=?`, phId)
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
