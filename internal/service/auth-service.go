package service

import (
	"context"
	"log"

	authaction "github.com/Fajar3108/online-course-be/internal/action/auth-action"
	useraction "github.com/Fajar3108/online-course-be/internal/action/user-action"
	authrequest "github.com/Fajar3108/online-course-be/internal/request/auth-request"
	errorhandler "github.com/Fajar3108/online-course-be/pkg/error-handler"
	"github.com/Fajar3108/online-course-be/pkg/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		DB: db,
	}
}

func (as *AuthService) Login(ctx context.Context, request *authrequest.LoginRequest) (jwToken string, refreshToken string, user *model.User, err error) {
	user = &model.User{}
	result := as.DB.WithContext(ctx).Where("email = ?", request.Email).First(user)

	if result.Error != nil {
		return "", "", nil, fiber.NewError(fiber.StatusUnauthorized, "Your email is not registered")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return "", "", nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
	}

	jwToken, refreshToken, err = authaction.CreateSessionAndTokens(ctx, user, as.DB)

	if err != nil {
		return "", "", nil, err
	}

	return jwToken, refreshToken, user, nil
}

func (as *AuthService) Register(ctx context.Context, request *authrequest.RegisterRequest) (string, string, *model.User, error) {
	var user *model.User
	var jwToken, refreshToken string

	err := as.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) (fiberErr error) {
		user, fiberErr = useraction.CreateNewUser(ctx, request, "user", tx)
		if fiberErr != nil {
			return fiberErr
		}

		jwToken, refreshToken, fiberErr = authaction.CreateSessionAndTokens(ctx, user, tx)

		if fiberErr != nil {
			return fiberErr
		}

		return nil
	})

	if err != nil {
		return "", "", nil, err
	}

	userName := user.Name
	userEmail := user.Email

	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic recovered in SendWelcomeEmail: %v", r)
			}
		}()

		authaction.SendWelcomeEmail(userName, userEmail)
	}()

	return jwToken, refreshToken, user, nil
}

func (as *AuthService) Logout(ctx context.Context, tokenJwt *jwt.Token) error {
	if err := as.DB.WithContext(ctx).Where("token = ?", tokenJwt.Raw).Delete(&model.UserSession{}).Error; err != nil {
		return errorhandler.GormErrorToFiberError(err)
	}

	return nil
}

func (as *AuthService) RefreshToken(ctx context.Context, request *authrequest.RefreshTokenRequest) (string, string, error) {
	userSession := &model.UserSession{}

	if err := as.DB.WithContext(ctx).Where("refresh_token = ?", request.RefreshToken).First(userSession).Error; err != nil {
		return "", "", errorhandler.GormErrorToFiberError(err)
	}

	user := &model.User{}
	if err := as.DB.WithContext(ctx).Where("id = ?", userSession.UserID).First(user).Error; err != nil {
		return "", "", errorhandler.GormErrorToFiberError(err)
	}

	jwToken, refreshToken, tokenExpired, refreshExpired, err := authaction.GenerateAuthToken(user)

	if err != nil {
		return "", "", fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	userSession.Token = jwToken
	userSession.TokenExpired = tokenExpired
	userSession.RefreshToken = refreshToken
	userSession.RefreshExpired = refreshExpired
	as.DB.WithContext(ctx).Save(userSession)

	return jwToken, refreshToken, nil
}
