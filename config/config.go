package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type ConfigStruct struct {
	App struct {
		Port string
	}
	Database struct {
		Host string
		Port string
		User string
		Pass string
		Name string
	}
	JWT struct {
		SecretKey         string
		Expiration        int
		RefreshExpiration int
	}
	SMTP struct {
		Host     string
		Port     int
		Username string
		Password string
		Sender   string
	}
	CookieSecretKey string
}

var (
	config *ConfigStruct
	once   sync.Once
)

func Config() *ConfigStruct {
	once.Do(func() {
		loadConfig()
	})

	return config
}

func loadConfig() {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	viper.AutomaticEnv()

	config = &ConfigStruct{}

	config.App.Port = viper.GetString("APP_PORT")

	config.Database.Host = viper.GetString("DB_HOST")
	config.Database.Port = viper.GetString("DB_PORT")
	config.Database.User = viper.GetString("DB_USER")
	config.Database.Pass = viper.GetString("DB_PASSWORD")
	config.Database.Name = viper.GetString("DB_NAME")

	config.JWT.SecretKey = viper.GetString("JWT_SECRET_KEY")
	config.JWT.Expiration = viper.GetInt("JWT_EXPIRATION_HOURS")
	config.JWT.RefreshExpiration = viper.GetInt("JWT_REFRESH_EXPIRATION_DAYS")

	config.SMTP.Host = viper.GetString("SMTP_HOST")
	config.SMTP.Port = viper.GetInt("SMTP_PORT")
	config.SMTP.Username = viper.GetString("SMTP_USERNAME")
	config.SMTP.Password = viper.GetString("SMTP_PASSWORD")
	config.SMTP.Sender = viper.GetString("SMTP_SENDER")

	config.CookieSecretKey = viper.GetString("COOKIE_SECRET_KEY")

}
