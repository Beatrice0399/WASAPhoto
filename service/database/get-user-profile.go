package database

import (
	"database/sql"
	"errors"
	"log"
)

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
		log.Print("errUser")
		return profile, ErrProfileDoesNotExist
	}

	follower, err := db.GetFollower(id)
	if err != nil {
		log.Print("errFolower")
		return profile, err
	}
	profile.Followers = len(follower)

	following, err := db.GetFollowing(id)
	if err != nil {
		log.Print("errFollowing")
		return profile, err
	}
	profile.Following = len(following)

	profile.NumberPhotos, _ = db.GetNumberPhotoUser(id)
	profile.Photos, _ = db.GetPhotoUser(id)

	return profile, err
}
