package database

func (db *appdbimpl) GetMyStream(myid int) ([]Photo, error) {

	var stream []Photo
	rows, err := db.c.Query(`SELECT p.* 	
							    FROM Photo p 
								JOIN (SELECT user FROM Follow WHERE followedBy=?) f ON p.user = f.user
								ORDER BY p.date DESC;`)
	for rows.Next() {
		var p Photo
		err = rows.Scan(&p.ID, &p.User, &p.Image, &p.Date, &p.Likes, &p.Comments)
		if err != nil {
			return stream, err
		}
		comments, err := db.GetPhotoComments(p.User)
		if err != nil {
			return stream, err
		}
		p.Comments = comments
		stream = append(stream, p)
	}

	return stream, err

	/*
		users, err := db.GetFollowing(myid)
		if err != nil {
			return stream, err
		}

		for _, value := range users {
			photos, err := db.GetPhotoUser(value.ID)
			if err != nil {
				return stream, err
			}
			stream = append(stream, photos...)
		}
	*/

	/*
		rows, err := db.c.Query(`SELECT p.* FROM Photo p JOIN (SELECT * FROM GetFollowing(?)) f ON p.user = f.id;`, myid)
		if err != nil {
			return stream, nil
		}

		for rows.Next() {
			var p Photo
			err = rows.Scan(&p.ID, &p.User, &p.Image, &p.Likes)
		}
		return stream, err
	*/
}
