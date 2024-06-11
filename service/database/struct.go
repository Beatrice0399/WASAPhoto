package database

import "time"

type Photo struct {
	ID       int       `json:"phid"`
	User     string    `json:"user"`
	Username string    `json:"username"`
	Date     time.Time `json:"date"`
	Likes    []Like    `json:"likes"`
	Comments []Comment `json:"comments"`
}

type User struct {
	Uid      int    `json:"uid"`
	Username string `json:"username"`
}

type Like struct {
	Uid int `json:"uid"`
}

type Comment struct {
	ID   int       `json:"id"`
	Uid  int       `json:"uid"`
	User string    `json:"user"`
	Text string    `json:"string"`
	Date time.Time `json:"date"`
}

type Profile struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Followers []User  `json:"followers"`
	Following []User  `json:"following"`
	Photos    []Photo `json:"photos"`
}
