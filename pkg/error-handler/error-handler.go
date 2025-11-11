package errorhandler

import (
	"errors"
	"log"

	"github.com/Fajar3108/online-course-be/pkg/helpers"
	"github.com/gofiber/fiber/v2"
)

func GlobalErrorHandler(ctx *fiber.Ctx, err error) error {
	res := helpers.NewResponseHelper(
		fiber.StatusInternalServerError,
		err.Error(),
		nil,
		nil,
		nil,
	)

	switch e := err.(type) {
	case *ValidationError:
		res.Code = fiber.StatusBadRequest
		res.Trace = e.Details
		return ctx.Status(res.Code).JSON(res)
	}

	var e *fiber.Error

	if errors.As(err, &e) {
		res.Code = e.Code
	}

	log.Printf("Error: %v\n", err)

	return ctx.Status(res.Code).JSON(res)
}
