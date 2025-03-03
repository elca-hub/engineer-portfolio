package email

import (
	"fmt"
	"net/smtp"
	"strings"
)

type Smtp struct{}

func NewSmtp() *Smtp {
	return &Smtp{}
}

func (s *Smtp) SendEmail(to []string, subject string, body string) error {
	config := NewSMTPConfig()

	smtpServer := fmt.Sprintf("%s:%s", config.smtpServer, config.smtpPort)

	msg := []byte(fmt.Sprintf("To: %s\nSubject: %s\n\n%s", strings.Join(to, ","), subject, body))

	return smtp.SendMail(smtpServer, nil, config.smtpUser, to, msg)
}
