package email

import (
	"bytes"
	"gopkg.in/gomail.v2"
	"html/template"
	"strconv"
)

type Smtp struct{}

func NewSmtp() *Smtp {
	return &Smtp{}
}

func (s *Smtp) SendEmail(to string, subject string, vars interface{}, files ...string) error {
	config := NewSMTPConfig()

	body, err := makeBody(vars, files...)
	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", config.smtpUser)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	portNum, _ := strconv.Atoi(config.smtpPort)

	println(body)
	d := gomail.Dialer{Host: config.smtpServer, Port: portNum}

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func makeBody(vars interface{}, files ...string) (string, error) {
	t, err := template.ParseFiles(files...)
	if err != nil {
		return "", err
	}
	buff := &bytes.Buffer{}

	if err := t.Execute(buff, vars); err != nil {
		return "", err
	}

	return buff.String(), nil
}
