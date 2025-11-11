package mail

import (
	"github.com/Fajar3108/online-course-be/config"
	"gopkg.in/gomail.v2"
)

func SendMail(receiver string, subject string, body string) error {
	mailer := gomail.NewMessage()

	mailer.SetHeader("From", config.Config().SMTP.Sender)
	mailer.SetHeader("To", receiver)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	dialer := gomail.NewDialer(
		config.Config().SMTP.Host,
		config.Config().SMTP.Port,
		config.Config().SMTP.Username,
		config.Config().SMTP.Password,
	)

	return dialer.DialAndSend(mailer)
}
