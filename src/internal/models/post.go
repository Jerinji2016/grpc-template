package models

import uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"

type Post struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Message   string
	CreatedAt string
	UserID    uuid.UUID `gorm:"index"`
	User      User      `gorm:"constraints:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
