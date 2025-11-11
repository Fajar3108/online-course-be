package authaction

import (
	"context"

	authrequest "github.com/Fajar3108/online-course-be/internal/request/auth-request"
	"github.com/Fajar3108/online-course-be/pkg/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateSessionAndTokens(ctx context.Context, user *model.User, db *gorm.DB) (string, string, error) {
	jwToken, refreshToken, tokenExpired, refreshExpired, tokenErr := GenerateAuthToken(user)

	if tokenErr != nil {
		return "", "", fiber.NewError(fiber.StatusInternalServerError, tokenErr.Error())
	}

	userSessionRequest := &authrequest.UserSessionRequest{
		UserID:         user.ID,
		Token:          jwToken,
		RefreshToken:   refreshToken,
		TokenExpired:   tokenExpired,
		RefreshExpired: refreshExpired,
	}

	if _, fiberErr := CreateNewUserSession(ctx, userSessionRequest, db); fiberErr != nil {
		return "", "", fiberErr
	}

	return jwToken, refreshToken, nil
}
