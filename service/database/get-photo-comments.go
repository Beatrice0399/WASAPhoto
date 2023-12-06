package database

import "database/sql"

func (db *appdbimpl) GetPhotoComments(phId int) (*sql.Rows, error) {
	//var comments []Comment
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
