package database

// Database function that removes all comments of the user (banned) from the photos of the user (myId) that banned him
func (db *appdbimpl) removeAllComments(myId int, banned int) error {
	res, err := db.c.Exec(`DELETE FROM Comment 
							WHERE id IN (
								SELECT c.id
								FROM Comment c
								JOIN Photo p ON p.user = ?
								WHERE c.user = ?
								);`, myId, banned)

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

// Database function that removes all likes of the user (banned) from the photos of the user (myId) that banned him
func (db *appdbimpl) removeAllLikes(myId int, banned int) error {
	res, err := db.c.Exec(`DELETE FROM Likes 
							WHERE uid = ? AND phid IN (
								SELECT p.id
								FROM Photo p 
								WHERE p.user = ?
								);`, myId, banned)

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

// Database function that checks if the given username already exists
func (db *appdbimpl) UsernameExist(name string) bool {
	row := db.c.QueryRow(`SELECT count(*) FROM user WHERE username=?`, name)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return false
	}
	if count > 0 {
		return true
	} else {
		return false
	}
}
