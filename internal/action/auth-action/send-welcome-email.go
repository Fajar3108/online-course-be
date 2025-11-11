package authaction

import (
	"bytes"
	_ "embed"
	"text/template"

	"github.com/Fajar3108/online-course-be/pkg/mail"
	"github.com/gofiber/fiber/v2/log"
)

//go:embed template/welcome.html
var welcomeEmailTemplate string

func SendWelcomeEmail(userName string, userEmail string) {
	tmpl, err := template.New("welcome").Parse(welcomeEmailTemplate)

	if err != nil {
		log.Errorf("Failed to parse embedded template: %v", err)
		return
	}

	var body bytes.Buffer

	data := map[string]string{
		"Name": userName,
	}

	if err := tmpl.Execute(&body, data); err != nil {
		log.Errorf("Failed to execute email template: %v", err)
		return
	}

	if err := mail.SendMail(
		userEmail,
		"Welcome to Mafi Course",
		body.String(),
	); err != nil {
		log.Errorf("Failed to send an email: %v", err)
		return
	}
}
