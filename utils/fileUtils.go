package utils

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/ProtonMail/gopenpgp/v2/helper"
	"github.com/halra/mailra/domain"
)

func ProcessFiles(files []*multipart.FileHeader, publicKey string) ([]domain.EncryptedAttachment, error) {
	var attachments []domain.EncryptedAttachment

	for _, file := range files {
		attachment, err := processFile(file, publicKey)
		if err != nil {
			return nil, err
		}
		attachments = append(attachments, attachment)
	}

	return attachments, nil
}

func processFile(file *multipart.FileHeader, publicKey string) (domain.EncryptedAttachment, error) {
	src, err := file.Open()
	if err != nil {
		return domain.EncryptedAttachment{}, fmt.Errorf("failed to open file: %w", err)
	}
	defer src.Close()

	fileContent, err := ioutil.ReadAll(src)
	if err != nil {
		return domain.EncryptedAttachment{}, fmt.Errorf("failed to read file: %w", err)
	}

	encryptedContent, err := encryptPGP(fileContent, publicKey)
	if err != nil {
		return domain.EncryptedAttachment{}, fmt.Errorf("failed to encrypt file: %w", err)
	}

	return domain.EncryptedAttachment{
		Filename: file.Filename,
		Content:  encryptedContent,
	}, nil
}

func encryptPGP(data []byte, publicKey string) ([]byte, error) {
	encryptedData, err := helper.EncryptBinaryMessageArmored(publicKey, data)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt message: %w", err)
	}
	return []byte(encryptedData), nil
}

func CreateTempFilePath(filename string) string {
	absolutePath, _ := filepath.Abs(fmt.Sprintf("./tmp/%d/%s.pgp", time.Now().UnixNano(), filename))
	return absolutePath
}
