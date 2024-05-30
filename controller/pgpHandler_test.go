package controller

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"github.com/gin-gonic/gin"
	"github.com/halra/mailra/test"
	"github.com/stretchr/testify/assert"
)

// TestMain sets up the test environment and runs the tests.
func TestMain(m *testing.M) {
	var stop = make(chan interface{})
	go test.StartMockSmtpServer(stop)
	time.Sleep(3 * time.Second)
	code := m.Run()
	close(stop)
	time.Sleep(3 * time.Second)
	os.Exit(code)
}

func TestPgpHandler(t *testing.T) {

	rsaKey, _ := crypto.GenerateKey("", "hi@example.com", "rsa", 2048)
	publicKey, _ := rsaKey.GetArmoredPublicKey()

	time.Sleep(3 * time.Second)
	// Set gin to test mode
	gin.SetMode(gin.TestMode)

	// Create a new gin router
	r := gin.Default()

	// Register the handler
	r.POST("/pgp", PgpHandler)

	// Create a buffer to hold the form data
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	// Add form fields
	writer.WriteField("from", "test@example.com")
	writer.WriteField("to", "test@example.com")
	writer.WriteField("subject", "Test Subject")
	writer.WriteField("smtPServer", "localhost")
	writer.WriteField("smtpPort", "30666")
	writer.WriteField("smtpPassword", "password")
	writer.WriteField("smtpUser", "user")
	writer.WriteField("method", "mime")
	writer.WriteField("bodyText", "This is a test email.")
	writer.WriteField("publicKey", publicKey)

	// Add a file
	part, err := writer.CreateFormFile("files", "testfile.txt")
	if err != nil {
		t.Fatalf("CreateFormFile error: %v", err)
	}
	part.Write([]byte("testdata"))

	part2, err := writer.CreateFormFile("files", "testfile.txt")
	if err != nil {
		t.Fatalf("CreateFormFile error: %v", err)
	}
	part2.Write([]byte("testdata"))

	writer.Close()

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodPost, "/pgp", body)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Create a new response recorder
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	assert.Equal(t, "Files uploaded and sent successfully.", w.Body.String())
}

func TestPgpHandlerInline(t *testing.T) {

	rsaKey, _ := crypto.GenerateKey("", "hi@example.com", "rsa", 2048)
	publicKey, _ := rsaKey.GetArmoredPublicKey()

	time.Sleep(3 * time.Second)
	// Set gin to test mode
	gin.SetMode(gin.TestMode)

	// Create a new gin router
	r := gin.Default()

	// Register the handler
	r.POST("/pgp", PgpHandler)

	// Create a buffer to hold the form data
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	// Add form fields
	writer.WriteField("from", "test@example.com")
	writer.WriteField("to", "test@example.com")
	writer.WriteField("subject", "Test Subject")
	writer.WriteField("smtPServer", "localhost")
	writer.WriteField("smtpPort", "30666")
	writer.WriteField("smtpPassword", "password")
	writer.WriteField("smtpUser", "user")
	writer.WriteField("method", "inline")
	writer.WriteField("bodyText", "This is a test email.")
	writer.WriteField("publicKey", publicKey)

	// Add a file
	part, err := writer.CreateFormFile("files", "testfile.txt")
	if err != nil {
		t.Fatalf("CreateFormFile error: %v", err)
	}
	part.Write([]byte("testdata"))

	part2, err := writer.CreateFormFile("files", "testfile.txt")
	if err != nil {
		t.Fatalf("CreateFormFile error: %v", err)
	}
	part2.Write([]byte("testdata"))

	writer.Close()

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodPost, "/pgp", body)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Create a new response recorder
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	assert.Equal(t, "Files uploaded and sent successfully.", w.Body.String())
}
