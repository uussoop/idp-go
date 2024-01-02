package database

type User struct {
	ID       int    `gorm:"primary_key"   json:"id"`
	Address  string `gorm:"unique"`
	Nonce    string
	Username string `gorm:"unique"        json:"email"`
	Verified bool   `gorm:"default:false" json:"verified"`
}
