package model

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"column:username;unique"`
	Password string `gorm:"column:password"`
}
