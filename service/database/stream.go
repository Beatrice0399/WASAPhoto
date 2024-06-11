package database

// Database function that returns the list of photos of followed
func (db *appdbimpl) GetMyStream(myid int) ([]Photo, error) {

	var stream []Photo
	rows, err := db.c.Query(`SELECT p.id, p.user, u.username, p.date	
							    FROM Photo p 
								JOIN (SELECT user FROM Follow WHERE followedBy=?) f ON p.user = f.user
								JOIN User u ON u.id = p.user
								ORDER BY p.date DESC;`, myid)
	if err != nil {
		return stream, err
	}
	defer rows.Close()
	for rows.Next() {
		var p Photo
		err = rows.Scan(&p.ID, &p.User, &p.Username, &p.Date)
		if err != nil {
			return stream, err
		}
		p.Likes, err = db.GetLikesPhoto(p.ID)
		if err != nil {
			return nil, err
		}
		p.Comments, err = db.GetPhotoComments(p.ID)
		if err != nil {
			return stream, err
		}
		stream = append(stream, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return stream, err
}
