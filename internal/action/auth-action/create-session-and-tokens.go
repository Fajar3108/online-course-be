package authaction

import (
	"context"

	authrequest "github.com/Fajar3108/online-course-be/internal/request/auth-request"
	errorhandler "github.com/Fajar3108/online-course-be/pkg/error-handler"
	"github.com/Fajar3108/online-course-be/pkg/model"
	"gorm.io/gorm"
)

func CreateSessionAndTokens(ctx context.Context, user *model.User, db *gorm.DB) (string, string, error) {
	jwToken, refreshToken, tokenExpired, refreshExpired, tokenErr := GenerateAuthToken(user)

	if tokenErr != nil {
		return "", "", errorhandler.GormErrorToFiberError(tokenErr)
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
