package database

func (db *appdbimpl) UnfollowUser(myId int, idProfile int) error {
	res, err := db.c.Exec(`DELETE FROM Follow WHERE user=? AND followedBy=?`, idProfile, myId)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		return ErrProfileDoesNotExist
	}
	return nil
}
