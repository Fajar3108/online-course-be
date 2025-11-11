package service

import (
	"context"

	"github.com/Fajar3108/online-course-be/internal/request"
	errorhandler "github.com/Fajar3108/online-course-be/pkg/error-handler"
	"github.com/Fajar3108/online-course-be/pkg/helpers"
	"github.com/Fajar3108/online-course-be/pkg/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CategoryService struct {
	db *gorm.DB
}

func NewCategoryService(db *gorm.DB) *CategoryService {
	return &CategoryService{
		db: db,
	}
}

func (cs *CategoryService) GetAll(ctx context.Context, page, limit int) (categories []model.Category, err error) {
	offset := (page - 1) * limit

	result := cs.db.WithContext(ctx).Order("name ASC").Offset(offset).Limit(limit).Find(&categories)

	if result.Error != nil {
		return nil, errorhandler.GormErrorToFiberError(result.Error)
	}

	return categories, nil
}

func (cs *CategoryService) GetBySlug(ctx context.Context, slug string) (category *model.Category, err error) {
	result := cs.db.WithContext(ctx).First(&category, "slug = ?", slug)

	if result.Error != nil {
		return nil, errorhandler.GormErrorToFiberError(result.Error)
	}

	return category, nil
}

func (cs *CategoryService) Store(ctx context.Context, categoryRequest *request.CategoryRequest) (category *model.Category, err error) {
	id, err := helpers.GenerateUUID()

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	category = &model.Category{
		ID:   id,
		Name: categoryRequest.Name,
		Slug: helpers.Slug(categoryRequest.Name),
	}

	if result := cs.db.WithContext(ctx).Create(category); result.Error != nil {
		return nil, errorhandler.GormErrorToFiberError(result.Error)
	}

	return category, nil
}

func (cs *CategoryService) Update(ctx context.Context, categoryRequest *request.CategoryRequest, slug string) (category *model.Category, err error) {
	category, err = cs.GetBySlug(ctx, slug)

	if err != nil {
		return nil, err
	}

	category.Name = categoryRequest.Name
	category.Slug = helpers.Slug(categoryRequest.Name)

	if result := cs.db.WithContext(ctx).Save(category); result.Error != nil {
		return nil, errorhandler.GormErrorToFiberError(result.Error)
	}

	return category, nil
}

func (cs *CategoryService) Destroy(ctx context.Context, slug string) error {
	category, err := cs.GetBySlug(ctx, slug)

	if err != nil {
		return err
	}

	if result := cs.db.WithContext(ctx).Delete(category); result.Error != nil {
		return errorhandler.GormErrorToFiberError(result.Error)
	}

	return nil
}
