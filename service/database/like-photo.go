package database

func (db *appdbimpl) LikePhoto(phId int, uid int) error {
	_, err := db.c.Exec(`INSERT INTO Likes (phId, uid) VALUES (?,?);`, phId, uid)
	if err != nil {
		return err
	}
	return err
}
