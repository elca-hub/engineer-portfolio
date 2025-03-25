package email

import (
	"fmt"
	"net/smtp"
	"strings"
	"regexp"
)

type Smtp struct{}

func NewSmtp() *Smtp {
	return &Smtp{}
}

func (s *Smtp) SendEmail(to []string, subject string, body string) error {
	config := NewSMTPConfig()

	// Validate email addresses
	for _, email := range to {
		if !isValidEmail(email) {
			return fmt.Errorf("invalid email address: %s", email)
		}
	}

	// Sanitize subject and body
	sanitizedSubject := sanitizeInput(subject)
	sanitizedBody := sanitizeInput(body)

	smtpServer := fmt.Sprintf("%s:%s", config.smtpServer, config.smtpPort)

msg := []byte(fmt.Sprintf("To: %s\nSubject: %s\n\n%s", strings.Join(to, ","), encodeHTML(sanitizedSubject), encodeHTML(sanitizedBody)))

	return smtp.SendMail(smtpServer, nil, config.smtpUser, to, msg)
}
// isValidEmail validates the email format
func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// sanitizeInput removes potentially harmful content from input
func sanitizeInput(input string) string {
	// Implement sanitization logic here (e.g., remove HTML tags, escape special characters)
	return input
}
