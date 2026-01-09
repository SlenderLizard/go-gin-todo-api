package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"` // unique user name
	Password string `json:"-"`        // hashed version, not visible in JSON
}
