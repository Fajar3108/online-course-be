package service

import (
	"context"
	"errors"

	"github.com/Fajar3108/online-course-be/database"
	"github.com/Fajar3108/online-course-be/internal/request"
	"github.com/Fajar3108/online-course-be/pkg/helpers"
	"github.com/Fajar3108/online-course-be/pkg/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CategoryService struct {
	db *gorm.DB
}

func NewCategoryService() *CategoryService {
	return &CategoryService{
		db: database.DB(),
	}
}

func (cs *CategoryService) GetAll(ctx context.Context, page, limit int) (categories *[]model.Category, err error) {
	offset := (page - 1) * limit

	result := cs.db.WithContext(ctx).Order("name ASC").Offset(offset).Limit(limit).Find(&categories)

	if result.Error != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}

	return categories, nil
}

func (cs *CategoryService) GetBySlug(ctx context.Context, slug string) (category *model.Category, err error) {
	result := cs.db.WithContext(ctx).First(&category, "slug = ?", slug)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fiber.NewError(fiber.StatusNotFound, "Category not found")
		}

		return nil, fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
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
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, fiber.NewError(fiber.StatusConflict, "Category with the same name already exists")
		}
		return nil, fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
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
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, fiber.NewError(fiber.StatusConflict, "Category with the same name already exists")
		}
		return nil, fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}

	return category, nil
}

func (cs *CategoryService) Destroy(ctx context.Context, slug string) error {
	category, err := cs.GetBySlug(ctx, slug)

	if err != nil {
		return err
	}

	if result := cs.db.WithContext(ctx).Delete(category); result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}

	return nil
}
