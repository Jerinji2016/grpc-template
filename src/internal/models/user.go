package models

import uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name     string
	Username string `gorm:"unique"`
	Password string
	Posts    []Post `gorm:"foreignKey:UserID"`
}
