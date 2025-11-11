package mail

import (
	"github.com/Fajar3108/online-course-be/config"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

func SendMail(receiver string, subject string, body string) error {
	mailer := gomail.NewMessage()

	mailer.SetHeader("From", viper.GetString(config.MailSender))
	mailer.SetHeader("To", receiver)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	dialer := gomail.NewDialer(
		viper.GetString(config.SMTPHost),
		viper.GetInt(config.SMTPPort),
		viper.GetString(config.MailUsername),
		viper.GetString(config.MailPassword),
	)

	return dialer.DialAndSend(mailer)
}
