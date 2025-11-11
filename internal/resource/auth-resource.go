package resource

import (
	"strings"

	"github.com/Fajar3108/online-course-be/pkg/model"
)

type UserAuth struct {
	ID     string `json:"id"`
	Name   string `json:"Name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
}

type AuthResource struct {
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
	User         *UserAuth `json:"user"`
}

func NewAuthResource(token string, refreshToken string, user *model.User, baseUrl string) *AuthResource {
	avatarURL := ""

	if user.Avatar != "" {
		avatarURL = strings.Join([]string{baseUrl, "storage", user.Avatar}, "/")
	}

	return &AuthResource{
		Token:        token,
		RefreshToken: refreshToken,
		User: &UserAuth{
			ID:     user.ID,
			Name:   user.Name,
			Email:  user.Email,
			Avatar: avatarURL,
		},
	}
}
