package database

func (db *appdbimpl) UncommentPhoto(cid int, phid int, uid int) error {
	res, err := db.c.Exec(`DELETE FROM Comment WHERE id=? AND user=? AND photo=?`, cid, uid, phid)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		return err
	}
	return nil
}
