package database

import "log"

func (db *appdbimpl) GetMyStream(myid int) ([]Photo, error) {

	var stream []Photo
	rows, err := db.c.Query(`SELECT p.id, u.username, p.image, p.date	
							    FROM Photo p 
								JOIN (SELECT user FROM Follow WHERE followedBy=?) f ON p.user = f.user
								JOIN User u ON u.id = p.user
								ORDER BY p.date DESC;`, myid)
	if err != nil {
		log.Println("ERRORE QUERY: ", err)
		return nil, err
	}
	for rows.Next() {
		var p Photo
		err = rows.Scan(&p.ID, &p.User, &p.Image, &p.Date)
		if err != nil {
			log.Println("ERR first scan: ", err)
			return stream, err
		}
		res, _ := db.GetLikesPhoto(p.ID)
		err = res.Scan(&p.Likes)

		com, _ := db.GetPhotoComments(p.ID)
		for exist := com.Next(); exist == true; exist = com.Next() {
			var c Comment
			err = com.Scan(&c.ID, &c.User, &c.Text, &c.Date)
			p.Comments = append(p.Comments, c)
		}
		stream = append(stream, p)
	}

	return stream, err
}
