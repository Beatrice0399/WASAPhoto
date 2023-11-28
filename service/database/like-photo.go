package database

func (db *appdbimpl) LikePhoto(phId uint64, uid uint64) error {
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
