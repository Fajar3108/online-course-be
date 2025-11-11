package authaction

import (
	"time"

	"github.com/Fajar3108/online-course-be/config"
	"github.com/Fajar3108/online-course-be/pkg/model"
	"github.com/Fajar3108/online-course-be/pkg/token"
)

func GenerateAuthToken(user *model.User) (string, string, *time.Time, *time.Time, error) {
	tokenDuration := time.Duration(config.Config().JWT.Expiration) * time.Hour
	tokenExpired := time.Now().Add(tokenDuration)
	jwToken, err := token.GenerateJWT(user, tokenExpired)

	if err != nil {
		return "", "", nil, nil, err
	}

	refreshDuration := time.Duration(config.Config().JWT.RefreshExpiration) * 24 * time.Hour
	refreshExpired := time.Now().Add(refreshDuration)
	refreshToken, err := token.GenerateJWT(user, refreshExpired)

	if err != nil {
		return "", "", nil, nil, err
	}

	return jwToken, refreshToken, &tokenExpired, &refreshExpired, nil
}
