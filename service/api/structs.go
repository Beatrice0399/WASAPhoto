package api

import "time"

const FORMAT_ERROR_IMG = "images must be jpeg or png"

type ErrMsgJSON struct {
	Message string `json:"message"` // Error message
}

type Username struct {
	Username string `json:"username"`
}

type User struct {
	Uid      int    `json:"uid"`
	Username string `json:"username"`
}

type Photo struct {
	Phid     int       `json:"uid"`
	User     int       `json:"user"`
	Username string    `json:"username"`
	Path     string    `json:"path"`
	Date     time.Time `json:"date"`
}

type CommentID struct {
	Cid int `json:"cid"`
}

type CommentText struct {
	Comment string `json:"comment"`
}
