package entities

// User struct
type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password []byte `json:"password"`
	IsAdmin  bool   `json:"isAdmin"`
}
