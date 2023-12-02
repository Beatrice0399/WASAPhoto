package database

import "log"

func (db *appdbimpl) GetUserProfile(id int) (Profile, error) {
	/*
		row := db.c.QueryRow(`SELECT * FROM Profile WHERE id=?`, id)
		var profile Profile
		err := row.Scan(&profile.ID, &profile.Name, &profile.Follower, &profile.Following, &profile.NumberPhotos, &profile.Photos)
	*/
	row := db.c.QueryRow(`SELECT * FROM User WHERE id=?`, id)
	var profile Profile
	err := row.Scan(&profile.ID, &profile.Name)
	log.Printf("GetUserProfile-GetUser: id: %d, name: %s\n", profile.ID, profile.Name)
	/*
		if err != nil {
			log.Print("errUser")
			return profile, ErrProfileDoesNotExist
		}
	*/

	follower, err := db.GetFollower(id)
	if err != nil {
		log.Print("errFolower")
		return profile, err
	}
	profile.Follower = len(follower)

	following, err := db.GetFollowing(id)
	if err != nil {
		log.Print("errFollowing")
		return profile, err
	}
	profile.Following = len(following)
	log.Printf("GetUserProfile. id: %d, name: %s, follower: %d, following: %d, pho: %d\n", profile.ID, profile.Name, profile.Follower, profile.Following, profile.NumberPhotos)

	return profile, err
}
