package models

type User struct {
	ID       string `gorm:"primaryKey"`
	Name     string
	Email    string `gorm:"unique"`
	Username string
	Password string
	Posts    []Post `gorm:"foreignKey:UserID"`
}
