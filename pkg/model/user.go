package model

import (
	"database/sql"
)

type User struct {
	ID        string       `json:"id" gorm:"primaryKey;not null"`
	Name      string       `json:"name" gorm:"not null;size:255"`
	Email     string       `json:"email" gorm:"unique;not null"`
	Role      string       `json:"role" gorm:"not null;size:50;default:'user'"`
	Avatar    string       `json:"avatar" gorm:"size:255"`
	Password  string       `json:"-" gorm:"not null;size:255"`
	CreatedAt sql.NullTime `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt sql.NullTime `json:"updated_at" gorm:"autoUpdateTime"`
}
