package helpers

import (
	"context"
	"math"

	"github.com/Fajar3108/online-course-be/database"
	"github.com/gofiber/fiber/v2"
)

type PaginationMeta struct {
	Page    int   `json:"page"`
	Limit   int   `json:"limit"`
	Total   int64 `json:"total"`
	MaxPage int   `json:"max_page"`
}

func NewPaginationMeta[T any](ctx context.Context, page, limit int) (*PaginationMeta, error) {
	var total int64
	db := database.DB()

	if err := db.WithContext(ctx).Model(new(T)).Count(&total).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	maxPage := 0
	if total > 0 && limit > 0 {
		maxPage = int(math.Ceil(float64(total) / float64(limit)))
	}

	return &PaginationMeta{
		Page:    page,
		Limit:   limit,
		Total:   total,
		MaxPage: maxPage,
	}, nil
}
