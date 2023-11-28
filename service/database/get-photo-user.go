package database

func (db *appdbimpl) GetPhotoUser(id uint64) ([]Photo, error) {
	var photos []Photo
	rows, err := db.c.Query(`SELECT * FROM Photo WHERE user=? ORDER BY date DESC`, id)
	if err != nil {
		return nil, err
	}

	defer func() { _ = rows.Close() }()

	for rows.Next() {

		var p Photo
		err = rows.Scan(&p.ID, &p.User, &p.Image, &p.Date, &p.Likes, &p.Comments)
		if err != nil {
			return nil, err
		}
		comments, err := db.GetPhotoComments(p.User)
		if err != nil {
			return nil, err
		}
		p.Comments = comments
		photos = append(photos, p)
	}
	return photos, err
}
