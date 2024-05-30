package utils

import (
	"fmt"
	"strconv"

	"gopkg.in/gomail.v2"

	"github.com/halra/mailra/domain"
)

func SendEmail(m *gomail.Message, mr domain.MailRequest) error {
	port, err := strconv.Atoi(mr.SmtpPort)
	if err != nil {
		return fmt.Errorf("invalid SMTP port: %w", err)
	}

	d := gomail.NewDialer(mr.SmtPServer, port, mr.SmtpUser, mr.SmtpPassword)
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
