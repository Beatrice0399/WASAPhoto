package database

import "time"

func (db *appdbimpl) UploadPhoto(id int, img []byte) (Photo, error) {
	var photo Photo
	photo.Image = img
	istante := time.Now()
	date := istante.Format("2006-01-02 15:04:05")

	res, err := db.c.Exec(`INSERT INTO Photo (user, image, date, likes) VALUES (?, ?, ?, NULL)`,
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
