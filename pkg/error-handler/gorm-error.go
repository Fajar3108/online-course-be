package errorhandler

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GormErrorToFiberError(err error) *fiber.Error {
	statusCode := fiber.StatusInternalServerError

	switch err {
	case gorm.ErrRecordNotFound:
		statusCode = fiber.StatusNotFound
	case gorm.ErrDuplicatedKey:
		statusCode = fiber.StatusConflict
	case gorm.ErrInvalidData:
		statusCode = fiber.StatusBadRequest
	}

	return fiber.NewError(statusCode, err.Error())
}
