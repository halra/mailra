package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"github.com/ProtonMail/gopenpgp/v2/helper"
	"github.com/halra/mailra/domain"
	"gopkg.in/gomail.v2"
)

func CheckKeyExpiration(key string) bool {

	// Decode the PGP key
	keyRing, err := crypto.NewKeyFromArmored(key)
	if err != nil {
		log.Fatalf("Error loading PGP key: %v", err)
	}
	// Get the primary key
	return keyRing.IsExpired()
}

func EncryptPGP(data []byte, publicKey string) ([]byte, error) {
	encryptedData, err := helper.EncryptBinaryMessageArmored(publicKey, data)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt message: %w", err)
	}
	return []byte(encryptedData), nil
}

func SendPgpEmailMIME(attachments []domain.EncryptedAttachment, mr domain.MailRequest) error {

	bodyText, err := helper.EncryptMessageArmored(mr.PublicKey, mr.BodyText)
	if err != nil {
		return fmt.Errorf("failed to EncryptMessageArmored : %w", err)
	}
	m := gomail.NewMessage()
	m.SetHeader("From", mr.From)
	m.SetHeader("To", mr.To)
	m.SetHeader("Subject", mr.Subject)
	m.SetBody("text/plain", bodyText)

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

func SendPgpEmailInline(attachments []domain.EncryptedAttachment, mr domain.MailRequest) error {
	m := gomail.NewMessage()
	m.SetHeader("From", mr.From)
	m.SetHeader("To", mr.To)
	m.SetHeader("Subject", mr.Subject)

	body := mr.BodyText + "\n\n"
	for _, attachment := range attachments {
		body += fmt.Sprintf("%s\n\n", attachment.Content)
	}
	m.SetBody("text/plain", body)

	return SendEmail(m, mr)
}
