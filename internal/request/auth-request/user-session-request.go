package authrequest

import "time"

type UserSessionRequest struct {
	UserID         string     `json:"user_id" validate:"required"`
	Token          string     `json:"token" gorm:"uniqueIndex" validate:"required"`
	RefreshToken   string     `json:"refresh_token" gorm:"uniqueIndex" validate:"required"`
	TokenExpired   *time.Time `json:"-" validate:"required"`
	RefreshExpired *time.Time `json:"-" validate:"required"`
}
