package utils

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/gomail.v2"
)

func SendEmail(to, subject, templateName string, data map[string]interface{}) error {
	from := os.Getenv("MAIL_FROM")
	fromName := os.Getenv("MAIL_FROM_NAME")

	host := os.Getenv("MAIL_HOST")
	port := os.Getenv("MAIL_PORT")
	username := os.Getenv("MAIL_USERNAME")
	password := os.Getenv("MAIL_PASSWORD")

	templatePath := filepath.Join("templates", templateName)
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	var bodyContent string
	buf := new(strings.Builder)
	if err := t.Execute(buf, data); err != nil {
		return err
	}
	bodyContent = buf.String()

	p := 2525
	fmt.Sscanf(port, "%d", &p)

	m := gomail.NewMessage()
	m.SetHeader("From", fmt.Sprintf("%s <%s>", fromName, from))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", bodyContent)

	d := gomail.NewDialer(host, p, username, password)
	return d.DialAndSend(m)
}
