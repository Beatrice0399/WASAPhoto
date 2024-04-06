package database

func (db *appdbimpl) GetAllProfiles() ([]Profile, error) {
	var profiles []Profile
	rows, err := db.c.Query(`SELECT * FROM User`)
	if err != nil {
		return profiles, nil
	}
	defer rows.Close()
	var id int
	var name string
	for rows.Next() {
		err = rows.Scan(&id, &name)
		// log.Printf("id: %d, name: %s\n", id, name)
		u, _ := db.GetUserProfile(id, 0)
		profiles = append(profiles, u)
		// log.Printf("id: %d, name: %s, follower: %d, following: %d, pho: %d\n", u.ID, u.Name, u.Follower, u.Following, &u.NumberPhotos)
		/*
			err = rows.Scan(&u.ID, &u.Name, &u.Follower, &u.Following, &u.NumberPhotos)
			if err != nil {
				return profiles, err
			}
			id, _ := db.GetId(u.Name)
			follower, _ := db.GetFollower(id)
			following, _ := db.GetFollowing(id)
			photos, _ := db.GetPhotoUser(id)
			u.Follower = len(follower)
			u.Following = len(following)
			u.NumberPhotos = len(photos)
			u.Photos = photos

			profiles = append(profiles, u)
		*/
	}
	if rows.Err() != nil {
		return nil, err
	}

	return profiles, err
}
