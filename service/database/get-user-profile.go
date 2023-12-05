package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) GetUserProfile(id int) (Profile, error) {
	ban := db.c.QueryRow(`SELECT id FROM Ban WHERE banned=? OR whoBan=?`, id, id)
	var ban_id int
	exist := ban.Scan(&ban_id)
	if !errors.Is(exist, sql.ErrNoRows) {
		row := db.c.QueryRow(`SELECT * FROM User WHERE id=?`, id)
		var profile Profile
		_ = row.Scan(&profile.ID, &profile.Name)
		return profile, nil
	}

	row := db.c.QueryRow(`SELECT * FROM User WHERE id=?`, id)
	var profile Profile
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
	//log.Printf("GetUserProfile. id: %d, name: %s, follower: %d, following: %d, pho: %d\n", profile.ID, profile.Name, profile.Follower, profile.Following, profile.NumberPhotos)

	return profile, err
}
