package database

func (db *appdbimpl) UnlikePhoto(phid int, myid int) error {
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
