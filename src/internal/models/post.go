package models

type Post struct {
	ID        string `gorm:"primaryKey"`
	Message   string
	CreatedAt string
	UserID    string `gorm:"index"`
}
