package authaction

import (
	"context"

	authrequest "github.com/Fajar3108/online-course-be/internal/request/auth-request"
	errorhandler "github.com/Fajar3108/online-course-be/pkg/error-handler"
	"github.com/Fajar3108/online-course-be/pkg/helpers"
	"github.com/Fajar3108/online-course-be/pkg/model"
	"gorm.io/gorm"
)

func CreateNewUserSession(ctx context.Context, request *authrequest.UserSessionRequest, db *gorm.DB) (*model.UserSession, error) {
	id, err := helpers.GenerateUUID()

	if err != nil {
		return nil, errorhandler.GormErrorToFiberError(err)
	}

	userSession := &model.UserSession{
		ID:             id,
		UserID:         request.UserID,
		Token:          request.Token,
		RefreshToken:   request.RefreshToken,
		TokenExpired:   request.TokenExpired,
		RefreshExpired: request.RefreshExpired,
	}

	result := db.WithContext(ctx).Create(userSession)

	if result.Error != nil {
		return nil, errorhandler.GormErrorToFiberError(result.Error)
	}

	return userSession, nil
}
