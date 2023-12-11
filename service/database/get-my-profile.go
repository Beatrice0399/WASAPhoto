package database

func (db *appdbimpl) GetMyProfile(myid int) (Profile, error) {
	var profile Profile
	row := db.c.QueryRow(`SELECT * FROM User WHERE id=?`, myid)

	err := row.Scan(&profile.ID, &profile.Name)
	if err != nil {
		return profile, ErrProfileDoesNotExist
	}

	follower, err := db.GetFollower(myid)
	if err != nil {
		return profile, err
	}
	profile.Followers = len(follower)

	following, err := db.GetFollowing(myid)
	if err != nil {
		return profile, err
	}
	profile.Following = len(following)

	profile.NumberPhotos, _ = db.GetNumberPhotoUser(myid)
	profile.Photos, _ = db.GetPhotoUser(myid)

	return profile, err
}
