package database

func (db *appdbimpl) SetMyUsername(id uint64, name string) error {
	res, err := db.c.Exec(`UPDATE Profile SET username=? WHERE id=?`, name, id)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then the fountain didn't exist
		return ErrProfileDoesNotExist
	}
	return nil
}
