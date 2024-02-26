package database

import "log"

func (db *appdbimpl) FollowUser(myId int, user string) error {
	Uid, err := db.GetId(user)
	if err != nil {
		return err
	}
	//Non posso seguire chi mi ha bannato
	exist, _ := db.IsBanned(Uid, myId)
	if exist == true {
		return ErrFollowUser
	}

	_ = db.UnbanUser(myId, user)
	//insert row
	_, err = db.c.Exec(`INSERT INTO Follow (user, followedBy) VALUES (?,?)`, Uid, myId)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) UnfollowUser(myId int, user string) error {
	idProfile, err := db.GetId(user)
	if err != nil {
		return err
	}

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

func (db *appdbimpl) GetFollower(id int) ([]User, error) {
	rows, err := db.c.Query(`SELECT f.followedBy
							FROM Follow f
							WHERE f.user=?`, id)
	if err != nil {
		return nil, err
	}

	var users []User
	for rows.Next() {
		var pid int
		err = rows.Scan(&pid)
		if err != nil {
			return nil, err
		}
		row := db.c.QueryRow(`SELECT * FROM User WHERE id=?`, pid)
		var u User
		err = row.Scan(&u.Uid, &u.Username)
		if err != nil {
			return nil, err
		}
		log.Printf("Function GetFollower. id: %d, name: %s\n", u.Uid, u.Username)
		users = append(users, u)
	}
	return users, err
}

func (db *appdbimpl) GetFollowing(followedBy int) ([]User, error) {
	rows, err := db.c.Query(`SELECT u.* FROM User u
							JOIN Follow f ON u.id=f.followedBy WHERE f.followedBy=?`, followedBy)
	if err != nil {
		return nil, err
	}

	var users []User
	for rows.Next() {
		var u User
		err = rows.Scan(&u.Uid, &u.Username)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, err
}
