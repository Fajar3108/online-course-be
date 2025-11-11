package model

import "database/sql"

type Category struct {
	ID        string       `json:"id" gorm:"primaryKey;not null"`
	Name      string       `json:"name" gorm:"not null;size:255"`
	Slug      string       `json:"slug" gorm:"unique;not null"`
	CreatedAt sql.NullTime `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt sql.NullTime `json:"updated_at" gorm:"autoUpdateTime"`
}
