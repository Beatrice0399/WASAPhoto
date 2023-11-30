package database

func (db *appdbimpl) LikePhoto(phId int, uid int) error {
	exist, err := db.c.Query(`SELECT * FROM Likes WHERE phId=? AND uid=?`, phId, uid)
	if err != nil {
		return err
	} else if exist != nil {
		return ErrAlreadyLiked
	}
	_, err = db.c.Exec(`INSERT INTO Likes (phId, uid) VALUES (?,?);`, phId, uid)
	if err != nil {
		return nil
	}
	return err
}
