package useraction

import (
	"context"

	authrequest "github.com/Fajar3108/online-course-be/internal/request/auth-request"
	errorhandler "github.com/Fajar3108/online-course-be/pkg/error-handler"
	file_storage "github.com/Fajar3108/online-course-be/pkg/file-storage"
	"github.com/Fajar3108/online-course-be/pkg/helpers"
	"github.com/Fajar3108/online-course-be/pkg/model"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateNewUser(ctx context.Context, request *authrequest.RegisterRequest, role string, db *gorm.DB) (*model.User, error) {
	id, err := helpers.GenerateUUID()

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	user := &model.User{
		ID:    id,
		Name:  request.Name,
		Email: request.Email,
		Role:  role,
	}

	if request.Avatar != nil {
		avatarPath, err := file_storage.Store(request.Avatar, "avatars", true)

		if err != nil {
			return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		user.Avatar = avatarPath
	}

	if result := db.WithContext(ctx).Where("email = ?", request.Email).First(user); result.Error == nil {
		return nil, fiber.NewError(fiber.StatusConflict, "email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, errorhandler.GormErrorToFiberError(err)
	}

	user.Password = string(hashedPassword)

	result := db.WithContext(ctx).Create(user)

	if result.Error != nil {
		return nil, errorhandler.GormErrorToFiberError(result.Error)
	}

	return user, nil
}
