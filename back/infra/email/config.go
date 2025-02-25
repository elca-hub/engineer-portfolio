package email

import "os"

type Config struct {
	smtpServer string
	smtpPort   string
	smtpUser   string
	smtpPass   string
}

func NewSMTPConfig() *Config {
	return &Config{
		smtpServer: os.Getenv("SMTP_SERVER"),
		smtpPort:   os.Getenv("SMTP_PORT"),
		smtpUser:   os.Getenv("SMTP_USER"),
		smtpPass:   os.Getenv("SMTP_PASS"),
	}
}
