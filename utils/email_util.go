package utils

import "fmt"

func SendEmail(to, subject, body string) error {
	// In production, integrate SMTP (e.g. SendGrid, Mailgun, etc.)
	fmt.Printf("To: %s\nSubject: %s\nBody: %s\n", to, subject, body)
	return nil
}
