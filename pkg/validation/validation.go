package validation

import (
	"sync"

	errorhandler "github.com/Fajar3108/online-course-be/pkg/error-handler"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/gofiber/fiber/v2"
)

var (
	validate *validator.Validate
	trans    ut.Translator
	once     sync.Once
)

func getValidator() (*validator.Validate, ut.Translator) {
	once.Do(func() {
		validate = validator.New(validator.WithRequiredStructEnabled())

		english := en.New()
		uni := ut.New(english, english)
		trans, _ = uni.GetTranslator("en")

		if err := en_translations.RegisterDefaultTranslations(validate, trans); err != nil {
			panic(err)
		}
	})

	return validate, trans
}

func fiberValidationError(err error) error {
	messages := make(map[string]string)

	_, trans = getValidator()

	ve, ok := err.(validator.ValidationErrors)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	for _, e := range ve {
		messages[e.Field()] = e.Translate(trans)
	}

	if len(messages) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Validation error")
	}

	return errorhandler.NewValidationError("Validation error", messages)
}

func Validate[T any](ctx *fiber.Ctx, request *T) (err error) {
	err = ctx.BodyParser(request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate, _ = getValidator()

	if errs := validate.Struct(request); errs != nil {
		return fiberValidationError(errs)
	}

	return nil
}
