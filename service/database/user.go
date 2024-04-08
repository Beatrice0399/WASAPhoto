package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) SearchUser(myId int, username string) ([]User, error) {
	rows, err := db.c.Query(`SELECT *
							FROM User
							WHERE username LIKE '%' || ? || '%' AND id NOT IN (SELECT whoBan FROM ban b JOIN user u ON u.id=b.whoBan WHERE (b.banned=? AND b.whoBan=u.id))`, username, myId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var u User
		err = rows.Scan(&u.Uid, &u.Username)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, err
}

func (db *appdbimpl) SetMyUsername(id int, name string) error {
	res, err := db.c.Exec(`UPDATE User SET username=? WHERE id=?`, name, id)
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

func (db *appdbimpl) GetUserProfile(id int, myId int) (Profile, error) {
	ban := db.c.QueryRow(`SELECT id FROM Ban WHERE banned=? AND whoBan=?`, id, myId)
	var ban_id int
	exist := ban.Scan(&ban_id)
	var profile Profile
	if !errors.Is(exist, sql.ErrNoRows) {
		return profile, ErrUserBanned
	}
	ban = db.c.QueryRow(`SELECT id FROM Ban WHERE banned=? AND whoBan=?`, myId, id)
	exist = ban.Scan(&ban_id)
	if !errors.Is(exist, sql.ErrNoRows) {
		return profile, ErrUserBannedYou
	}

	row := db.c.QueryRow(`SELECT * FROM User WHERE id=?`, id)

	err := row.Scan(&profile.ID, &profile.Name)
	if err != nil {
		// log.Print("errUser")
		return profile, ErrProfileDoesNotExist
	}

	follower, err := db.GetFollower(id)
	if err != nil {
		// log.Print("errFolower")
		return profile, err
	}
	profile.Followers = follower

	following, err := db.GetFollowing(id)
	if err != nil {
		// log.Print("errFollowing")
		return profile, err
	}
	profile.Following = following

	profile.NumberPhotos, _ = db.GetNumberPhotoUser(id)
	profile.Photos, _ = db.GetPhotoUser(id)

	return profile, err
}
