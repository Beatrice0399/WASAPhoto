package database

func (db *appdbimpl) DeletePhoto(pid int, myid int) error {
	res, err := db.c.Exec(`DELETE FROM Photo WHERE id=? AND user=?`, pid, myid)
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
