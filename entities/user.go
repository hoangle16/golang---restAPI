package entities

// User struct
type User struct {
	ID       uint   `json:"id"`
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"isAdmin"`
}
