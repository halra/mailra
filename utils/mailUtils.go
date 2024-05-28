package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"gopkg.in/gomail.v2"

	"github.com/halra/mailra/domain"
)

func SendEmailMIME(attachments []domain.EncryptedAttachment, mr domain.MailRequest) error {
	m := gomail.NewMessage()
	m.SetHeader("From", mr.From)
	m.SetHeader("To", mr.To)
	m.SetHeader("Subject", mr.Subject)
	m.SetBody("text/plain", mr.BodyText)

	for _, attachment := range attachments {
		fp := CreateTempFilePath(attachment.Filename)
		if err := os.MkdirAll(filepath.Dir(fp), os.ModeTemporary); err != nil {
			return fmt.Errorf("failed to create directories for file: %w", err)
		}

		file, err := os.Create(fp)
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}

		defer func(fp string) {
			file.Close()
			os.RemoveAll(filepath.Dir(fp))
		}(fp)

		if _, err = file.Write(attachment.Content); err != nil {
			return fmt.Errorf("failed to write to file: %w", err)
		}

		m.Attach(fp)
	}

	return SendEmail(m, mr)
}

func SendEmailInline(attachments []domain.EncryptedAttachment, mr domain.MailRequest) error {
	m := gomail.NewMessage()
	m.SetHeader("From", mr.From)
	m.SetHeader("To", mr.To)
	m.SetHeader("Subject", mr.Subject)

	body := mr.BodyText + "\n\nPlease find the encrypted files below:\n\n"
	for _, attachment := range attachments {
		body += fmt.Sprintf("-----BEGIN PGP MESSAGE-----\n\n%s\n-----END PGP MESSAGE-----\n\n", attachment.Content)
	}

	m.SetBody("text/plain", body)

	return SendEmail(m, mr)
}

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
