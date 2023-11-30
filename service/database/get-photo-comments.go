package database

func (db *appdbimpl) GetPhotoComments(phId int) ([]Comment, error) {
	var comments []Comment
	rows, err := db.c.Query(`SELECT * FROM Comment WHERE photo=?`, phId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var c Comment
		err = rows.Scan(&c.ID, &c.User, &c.Text, &c.Date)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	return comments, err
}
