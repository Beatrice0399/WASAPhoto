package database

import "log"

func (db *appdbimpl) removeAllComments(myId int, banned int) error {

	/*
		res, err := db.c.Exec(`UPDATE Comment SET visible=0
								WHERE id IN (
									SELECT c.id
									FROM Comment c
									JOIN Photo p ON p.user = ?
									WHERE c.user = ?
									);`, myId, idProfile)
	*/
	res, err := db.c.Exec(`DELETE FROM Comment 
							WHERE id IN (
								SELECT c.id
								FROM Comment c
								JOIN Photo p ON p.user = ?
								WHERE c.user = ?
								);`, myId, banned)

	if err != nil {
		log.Println(err)
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

func (db *appdbimpl) removeAllLikes(myId int, banned int) error {
	res, err := db.c.Exec(`DELETE FROM Likes 
							WHERE id IN (
								SELECT c.id
								FROM Comment c
								JOIN Photo p ON p.user = ?
								WHERE c.user = ?
								);`, myId, banned)

	if err != nil {
		log.Println(err)
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
