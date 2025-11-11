package model

import (
	"database/sql"
	"time"
)

type UserSession struct {
	ID             string `json:"id" gorm:"primaryKey;not null"`
	CreatedAt      sql.NullTime
	UpdatedAt      sql.NullTime
	UserID         string     `json:"user_id"`
	Token          string     `json:"token" gorm:"size:512"`
	RefreshToken   string     `json:"refresh_token" gorm:"size:512"`
	TokenExpired   *time.Time `json:"-"`
	RefreshExpired *time.Time `json:"-"`
}
