package database

func (db *appdbimpl) GetMyStream(myid int) ([]Photo, error) {

	var stream []Photo
	rows, err := db.c.Query(`SELECT p.id, u.username, p.image_path, p.date	
							    FROM Photo p 
								JOIN (SELECT user FROM Follow WHERE followedBy=?) f ON p.user = f.user
								JOIN User u ON u.id = p.user
								ORDER BY p.date DESC;`, myid)
	if err != nil {
		return stream, err
	}
	for rows.Next() {
		var p Photo
		err = rows.Scan(&p.ID, &p.User, &p.Path, &p.Date)
		if err != nil {
			return stream, err
		}
		res, err := db.GetLikesPhoto(p.ID)
		for res.Next() {
			var u User
			err = res.Scan(&u.Uid, &u.Username)
			if err != nil {
				return stream, err
			}
			p.Likes = append(p.Likes, u)
		}

		com, err := db.GetPhotoComments(p.ID)
		if err != nil {
			return stream, err
		}
		for exist := com.Next(); exist == true; exist = com.Next() {
			var c Comment
			err = com.Scan(&c.ID, &c.User, &c.Text, &c.Date)
			if err != nil {
				return stream, err
			}
			p.Comments = append(p.Comments, c)
		}
		stream = append(stream, p)
	}

	return stream, err
}
