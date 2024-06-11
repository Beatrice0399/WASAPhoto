package database

// Database function that allows an user (bid) to ban another one (uid)
func (db *appdbimpl) BanUser(bid int, uid int) error {
	_ = db.UnfollowUser(bid, uid)
	_ = db.removeAllComments(bid, uid)
	_ = db.removeAllComments(uid, bid)
	_ = db.removeAllLikes(bid, uid)
	_ = db.removeAllLikes(uid, bid)
	_ = db.UnfollowUser(uid, bid)

	_, err := db.c.Exec(`INSERT INTO Ban (banned, whoBan) VALUES (?,?)`, uid, bid)
	if err != nil {
		return err
	}

	return nil
}

// Database function that allows an user (bid) to unban another one (uid)
func (db *appdbimpl) UnbanUser(bid int, uid int) error {
	res, err := db.c.Exec(`DELETE FROM Ban WHERE banned=? AND whoBan=?`, uid, bid)
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
