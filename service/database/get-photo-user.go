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

		res, err := db.GetLikesPhoto(p.ID)
		if err != nil {
			return photos, err
		}
		defer res.Close()
		for res.Next() {
			var u User
			err = res.Scan(&u.Uid, &u.Username)
			if err != nil {
				return photos, err
			}
			p.Likes = append(p.Likes, u)
		}
		if res.Err() != nil {
			return nil, err
		}
		com, err := db.GetPhotoComments(p.ID)
		if err != nil {
			return photos, err
		}
		defer com.Close()
		for exist := com.Next(); exist == true; exist = com.Next() {
			var c Comment
			err = com.Scan(&c.ID, &c.User, &c.Text, &c.Date)
			if err != nil {
				return photos, err
			}
			p.Comments = append(p.Comments, c)
		}
		if com.Err() != nil {
			return nil, err
		}
		photos = append(photos, p)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return photos, err
}
