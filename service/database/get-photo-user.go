package database

func (db *appdbimpl) GetPhotoUser(id int) ([]Photo, error) {
	var photos []Photo
	rows, err := db.c.Query(`SELECT * FROM Photo WHERE user=? ORDER BY date DESC`, id)
	if err != nil {
		return photos, err
	}

	defer rows.Close()
	for rows.Next() {
		var p Photo
		err = rows.Scan(&p.ID, &p.User, &p.Path, &p.Date)
		if err != nil {
			return photos, err
		}

		p.Likes, err = db.GetLikesPhoto(p.ID)
		if err != nil {
			return photos, err
		}
		p.Comments, err = db.GetPhotoComments(p.ID)
		if err != nil {
			return photos, err
		}
		photos = append(photos, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return photos, err
}
