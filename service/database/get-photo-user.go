package database

func (db *appdbimpl) GetPhotoUser(id int) ([]Photo, error) {
	var photos []Photo
	rows, err := db.c.Query(`SELECT * FROM Photo WHERE user=? ORDER BY date DESC`, id)
	if err != nil {
		return nil, err
	}

	// defer func() { _ = rows.Close() }()

	for rows.Next() {

		var p Photo
		err = rows.Scan(&p.ID, &p.User, &p.Path, &p.Date)
		if err != nil {
			return nil, err
		}

		res, err := db.GetLikesPhoto(p.ID)
		for res.Next() {
			var u User
			err = res.Scan(&u.Uid, &u.Username)
			if err != nil {
				return nil, err
			}
			p.Likes = append(p.Likes, u)
		}

		com, _ := db.GetPhotoComments(p.ID)
		// var comments *[]Comment
		for exist := com.Next(); exist == true; exist = com.Next() {
			var c Comment
			err = com.Scan(&c.ID, &c.User, &c.Text, &c.Date)
			// *comments = append(*comments, *c)
			p.Comments = append(p.Comments, c)
		}
		photos = append(photos, p)
	}

	return photos, err
}
