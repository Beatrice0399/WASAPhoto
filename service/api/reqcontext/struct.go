package api

//import "github.com/Beatrice0399/WASAPhoto/service/database"

type Photo struct {
	ID       uint64    `json:"id"`
	User     string    `json:"User"`
	Image    []byte    `json:"image"`
	Likes    int       `json:"likes"`
	Comments []Comment `json:"comments"`
}

type User struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	ID   uint64 `json:"id"`
	User string `json:"user"`
	Text string `json:"string"`
}

type Stream struct {
	Photos []Photo `json:"photos"`
}

type Profile struct {
	ID           uint64 `json:"id"`
	Name         string `json:"name"`
	Following    int    `json:"following"`
	Follower     int    `json:"follower"`
	NumberPhotos int    `json:"numberPhotos"`
	Photos       Stream `json:"photos"`
}
