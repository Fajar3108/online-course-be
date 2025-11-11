package controller

import (
	"github.com/Fajar3108/online-course-be/internal/request"
	"github.com/Fajar3108/online-course-be/internal/resource"
	"github.com/Fajar3108/online-course-be/internal/service"
	"github.com/Fajar3108/online-course-be/pkg/helpers"
	"github.com/Fajar3108/online-course-be/pkg/model"
	"github.com/Fajar3108/online-course-be/pkg/validation"
	"github.com/gofiber/fiber/v2"
)

type CategoryController struct {
	service *service.CategoryService
}

func NewCategoryController(service *service.CategoryService) *CategoryController {
	return &CategoryController{
		service: service,
	}
}

func (cc *CategoryController) Index(ctx *fiber.Ctx) error {
	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 10)

	userContext := ctx.UserContext()
	categories, err := cc.service.GetAll(userContext, page, limit)

	if err != nil {
		return err
	}

	meta, err := helpers.NewPaginationMeta[model.Category](userContext, page, limit)

	if err != nil {
		return err
	}

	res := helpers.NewResponseHelper(
		fiber.StatusOK,
		"Categories retrieved successfully",
		helpers.NewResourceCollection(categories, resource.NewCategoryResource),
		meta,
		nil,
	)

	return ctx.Status(res.Code).JSON(res)
}

func (cc *CategoryController) Show(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")

	category, err := cc.service.GetBySlug(ctx.UserContext(), slug)

	if err != nil {
		return err
	}

	res := helpers.NewResponseHelper(
		fiber.StatusOK,
		"Category retrieved successfully",
		resource.NewCategoryResource(category),
		nil,
		nil,
	)

	return ctx.Status(res.Code).JSON(res)
}

func (cc *CategoryController) Store(ctx *fiber.Ctx) error {
	req := &request.CategoryRequest{}

	if err := validation.Validate(ctx, req); err != nil {
		return err
	}

	category, err := cc.service.Store(ctx.UserContext(), req)

	if err != nil {
		return err
	}

	res := helpers.NewResponseHelper(
		fiber.StatusOK,
		"Category created successfully",
		resource.NewCategoryResource(category),
		nil,
		nil,
	)

	return ctx.Status(res.Code).JSON(res)
}

func (cc *CategoryController) Update(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")

	req := &request.CategoryRequest{}

	if err := validation.Validate(ctx, req); err != nil {
		return err
	}

	category, err := cc.service.Update(ctx.UserContext(), req, slug)

	if err != nil {
		return err
	}

	res := helpers.NewResponseHelper(
		fiber.StatusOK,
		"Category updated successfully",
		resource.NewCategoryResource(category),
		nil,
		nil,
	)

	return ctx.Status(res.Code).JSON(res)
}

func (cc *CategoryController) Destroy(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")

	err := cc.service.Destroy(ctx.UserContext(), slug)

	if err != nil {
		return err
	}

	res := helpers.NewResponseHelper(
		fiber.StatusOK,
		"Category deleted successfully",
		nil,
		nil,
		nil,
	)

	return ctx.Status(res.Code).JSON(res)
}
