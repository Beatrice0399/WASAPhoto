package database

// Database funtion that add a like to the given photo
func (db *appdbimpl) LikePhoto(phId int, uid int) error {
	_, err := db.c.Exec(`INSERT INTO Likes (phId, uid) VALUES (?,?);`, phId, uid)
	if err != nil {
		return err
	}
	return err
}

// Database function that allows an user to remove the like from the photo
func (db *appdbimpl) UnlikePhoto(phid int, myid int, lid int) error {
	if myid != lid {
		return ErrLike
	}
	res, err := db.c.Exec(`DELETE FROM Likes WHERE phId=? AND uid=?`, phid, myid)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		return ErrPhoto
	}
	return nil
}

// Databse funtion that retuns ther list of the users who like the photo
func (db *appdbimpl) GetLikesPhoto(phid int) ([]Like, error) {
	rows, err := db.c.Query(`SELECT u.id FROM Likes l JOIN user u ON u.id = l.uid WHERE phId=?`, phid)
	if err != nil {
		return nil, err
	}
	var likes []Like
	defer rows.Close()
	for rows.Next() {
		var l Like
		err = rows.Scan(&l.Uid)
		if err != nil {
			return nil, err
		}
		likes = append(likes, l)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return likes, nil
}
