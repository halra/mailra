package utils

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"github.com/halra/mailra/domain"
	"github.com/halra/mailra/test"
	"github.com/stretchr/testify/assert"
	"gopkg.in/gomail.v2"
)

var randomPort = rand.Intn(38000-31000+1) + 31000

// TestMain sets up the test environment and runs the tests.
func TestMain(m *testing.M) {
	var stop = make(chan interface{})
	go test.StartMockSmtpServer(stop, fmt.Sprintf(":%d", randomPort))
	time.Sleep(3 * time.Second)
	code := m.Run()
	close(stop)
	time.Sleep(3 * time.Second)
	os.Exit(code)
}

func TestSendEmailMIME(t *testing.T) {

	// Create a test PGP public key for encryption
	rsaKey, _ := crypto.GenerateKey("", "hi@example.com", "rsa", 2048)
	publicKey, err := rsaKey.GetArmoredPublicKey()
	assert.NoError(t, err)

	// Set up test data
	mr := domain.MailRequest{
		From:         "sender@example.com",
		To:           "recipient@example.com",
		Subject:      "Test Subject",
		SmtPServer:   "localhost",
		SmtpPort:     fmt.Sprintf("%d", randomPort),
		SmtpPassword: "password",
		SmtpUser:     "user",
		Method:       "mime",
		BodyText:     "This is a test email.",
		PublicKey:    publicKey,
	}

	attachments := []domain.EncryptedAttachment{
		{
			Filename: "testfile.txt",
			Content:  []byte("this is a test file"),
		},
	}

	// Call the function under test
	err = SendPgpEmailMIME(attachments, mr)
	assert.NoError(t, err)

}

func TestSendEmailInline(t *testing.T) {

	// Set up test data
	mr := domain.MailRequest{
		From:         "sender@example.com",
		To:           "recipient@example.com",
		Subject:      "Test Subject",
		SmtPServer:   "localhost",
		SmtpPort:     fmt.Sprintf("%d", randomPort),
		SmtpPassword: "password",
		SmtpUser:     "user",
		Method:       "inline",
		BodyText:     "This is a test email.",
		PublicKey:    "testPublicKey",
	}

	attachments := []domain.EncryptedAttachment{
		{
			Filename: "testfile.txt",
			Content:  []byte("this is a test file"),
		},
	}

	// Call the function under test
	err := SendPgpEmailInline(attachments, mr)
	assert.NoError(t, err)

	// Assert that the mock dialer was called
}

func TestSendEmail(t *testing.T) {

	// Create a new gomail message
	m := gomail.NewMessage()
	m.SetHeader("From", "sender@example.com")
	m.SetHeader("To", "recipient@example.com")
	m.SetHeader("Subject", "Test Subject")
	m.SetBody("text/plain", "This is a test email.")

	// Set up test data
	mr := domain.MailRequest{
		From:         "sender@example.com",
		To:           "recipient@example.com",
		Subject:      "Test Subject",
		SmtPServer:   "localhost",
		SmtpPort:     fmt.Sprintf("%d", randomPort),
		SmtpPassword: "password",
		SmtpUser:     "user",
	}

	// Call the function under test
	err := SendEmail(m, mr)
	assert.NoError(t, err)

}
