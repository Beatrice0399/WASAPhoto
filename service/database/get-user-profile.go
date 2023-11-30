package database

func (db *appdbimpl) GetUserProfile(id int) (Profile, error) {
	row := db.c.QueryRow(`SELECT * FROM Profile WHERE id=?`, id)

	var profile Profile
	err := row.Scan(&profile.ID, &profile.Name, &profile.Follower, &profile.Following, &profile.NumberPhotos, &profile.Photos)

	if err != nil {
		return profile, ErrProfileDoesNotExist
	}
	follower, err := db.GetFollower(id)
	if err != nil {
		return profile, err
	}
	profile.Follower = len(follower)

	following, err := db.GetFollowing(id)
	if err != nil {
		return profile, err
	}
	profile.Following = len(following)

	return profile, err
}
