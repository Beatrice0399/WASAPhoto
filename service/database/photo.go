package database

import (
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
	row := db.c.QueryRow(`SELECT p.id, p.user, u.username, p.image_path, p.date
							FROM Photo p JOIN User u ON u.id=p.user
							WHERE p.id = ?`, phId)

	err := row.Scan(&photo.ID, &photo.User, &photo.Username, &photo.Path, &photo.Date)
	if err != nil {
		return photo, err
	}
	photo.Likes, err = db.GetLikesPhoto(photo.ID)
	if err != nil {
		return photo, err
	}
	photo.Comments, err = db.GetPhotoComments(photo.ID)
	if err != nil {
		return photo, err
	}
	return photo, nil
}

func (db *appdbimpl) GetPhotoComments(phId int) ([]Comment, error) {
	var comments []Comment
	rows, err := db.c.Query(`SELECT c.id, c.user, u.username, c.string, c.date FROM Comment c
								JOIN User u ON c.user=u.id WHERE photo=?`, phId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c Comment
		err = rows.Scan(&c.ID, &c.Uid, &c.User, &c.Text, &c.Date)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return comments, err
}
