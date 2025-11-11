package controller

import (
	authrequest "github.com/Fajar3108/online-course-be/internal/request/auth-request"
	"github.com/Fajar3108/online-course-be/internal/resource"
	"github.com/Fajar3108/online-course-be/internal/service"
	"github.com/Fajar3108/online-course-be/pkg/helpers"
	"github.com/Fajar3108/online-course-be/pkg/validation"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthController struct {
	service *service.AuthService
}

func NewAuthController(service *service.AuthService) *AuthController {
	return &AuthController{
		service: service,
	}
}

func (ac *AuthController) Login(ctx *fiber.Ctx) error {
	req := &authrequest.LoginRequest{}

	if err := validation.Validate(ctx, req); err != nil {
		return err
	}

	token, refreshToken, user, err := ac.service.Login(ctx.UserContext(), req)

	if err != nil {
		return err
	}

	data := resource.NewAuthResource(token, refreshToken, user, ctx.BaseURL())

	res := helpers.NewResponseHelper(
		fiber.StatusOK,
		"Login successful",
		data,
		nil,
		nil,
	)

	return ctx.Status(res.Code).JSON(res)
}

func (ac *AuthController) Register(ctx *fiber.Ctx) error {
	req := &authrequest.RegisterRequest{}

	if err := validation.Validate(ctx, req); err != nil {
		return err
	}

	if avatar, err := ctx.FormFile("avatar"); err != nil {
		req.Avatar = nil
	} else {
		req.Avatar = avatar
	}

	tokenJWT, refreshToken, user, err := ac.service.Register(ctx.UserContext(), req)

	if err != nil {
		return err
	}

	data := resource.NewAuthResource(tokenJWT, refreshToken, user, ctx.BaseURL())

	res := helpers.NewResponseHelper(
		fiber.StatusOK,
		"Registration successful",
		data,
		nil,
		nil,
	)

	return ctx.Status(res.Code).JSON(res)
}

func (ac *AuthController) Logout(ctx *fiber.Ctx) error {
	tokenJwt := ctx.Locals("user").(*jwt.Token)

	err := ac.service.Logout(ctx.UserContext(), tokenJwt)

	if err != nil {
		return err
	}

	res := helpers.NewResponseHelper(
		fiber.StatusOK,
		"Logout successful",
		nil,
		nil,
		nil,
	)

	return ctx.Status(res.Code).JSON(res)
}

func (ac *AuthController) RefreshToken(ctx *fiber.Ctx) error {
	req := &authrequest.RefreshTokenRequest{}

	if err := validation.Validate(ctx, req); err != nil {
		return err
	}

	newToken, newRefreshToken, err := ac.service.RefreshToken(ctx.UserContext(), req)

	if err != nil {
		return err
	}

	res := helpers.NewResponseHelper(
		fiber.StatusOK,
		"Token refreshed successfuly",
		map[string]string{
			"token":         newToken,
			"refresh_token": newRefreshToken,
		},
		nil,
		nil,
	)

	return ctx.Status(res.Code).JSON(res)
}
