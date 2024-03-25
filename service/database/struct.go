package database

import "time"

type Photo struct {
	ID       int       `json:"id"`
	User     string    `json:"User"`
	Image    []byte    `json:"image"`
	Date     time.Time `json:"date"`
	Likes    []User    `json:"likes"`
	Comments []Comment `json:"comments"`
}

type User struct {
	Uid      int    `json:"uid"`
	Username string `json:"username"`
}

type Comment struct {
	ID   int       `json:"id"`
	User string    `json:"user"`
	Text string    `json:"string"`
	Date time.Time `json:"date"`
}

type Profile struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Followers    []User  `json:"followers"`
	Following    []User  `json:"following"`
	NumberPhotos int     `json:"numberPhotos"`
	Photos       []Photo `json:"photos"`
}
