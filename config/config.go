package config

import (
	"log"

	"github.com/spf13/viper"
)

const (
	JwtSecretKey             = "JWT_SECRET_KEY"
	JwtExpirationHours       = "JWT_EXPIRATION_HOURS"
	JwtRefreshExpirationDays = "JWT_REFRESH_EXPIRATION_DAYS"

	DbHost     = "DB_HOST"
	DbPort     = "DB_PORT"
	DbUser     = "DB_USER"
	DbPassword = "DB_PASSWORD"
	DbName     = "DB_NAME"
	AppPort    = "APP_PORT"

	CookieSecretKey = "COOKIE_SECRET_KEY"

	SMTPHost     = "MAIL_HOST"
	SMTPPort     = "MAIL_PORT"
	MailUsername = "MAIL_USERNAME"
	MailPassword = "MAIL_PASSWORD"
	MailSender   = "MAIL_SENDER"
)

func InitConfig() {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	viper.AutomaticEnv()
}
