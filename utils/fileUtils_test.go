package utils

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"testing"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"github.com/stretchr/testify/assert"
)

// Helper function to create a temporary file and return its FileHeader.
func createMultipartFileHeader(t *testing.T, content string) (*multipart.FileHeader, *os.File) {
	// Create a temporary file with the given content.
	tempFile, err := ioutil.TempFile("", "testfile")
	if err != nil {
		t.Fatal(err)
	}

	_, err = tempFile.WriteString(content)
	if err != nil {
		t.Fatal(err)
	}

	_, err = tempFile.Seek(0, 0)
	if err != nil {
		t.Fatal(err)
	}

	// Create a multipart form file header.
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	part, err := writer.CreateFormFile("file", tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	_, err = io.Copy(part, tempFile)
	if err != nil {
		t.Fatal(err)
	}

	writer.Close()

	// Parse the multipart form to get the FileHeader.
	req := &http.Request{
		Header: make(http.Header),
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Body = ioutil.NopCloser(&buffer)
	err = req.ParseMultipartForm(10 << 20) // MaxMemory 10 MB
	if err != nil {
		t.Fatal(err)
	}

	return req.MultipartForm.File["file"][0], tempFile
}

func TestProcessFiles(t *testing.T) {
	// Create a test PGP public key for encryption
	rsaKey, _ := crypto.GenerateKey("", "hi@go.com", "rsa", 2048)
	publicKey, err := rsaKey.GetArmoredPublicKey()
	if err != nil {
		t.Fatal(err)
	}

	// Create a multipart file header with test content
	fileHeader, tempFile := createMultipartFileHeader(t, "this is a test file")
	defer os.Remove(tempFile.Name()) // Clean up the temporary file

	// Test ProcessFiles function with valid file
	files := []*multipart.FileHeader{fileHeader}
	attachments, err := ProcessFiles(files, publicKey)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(attachments))
	assert.Equal(t, fileHeader.Filename, attachments[0].Filename)
	assert.NotEmpty(t, attachments[0].Content)

	h := textproto.MIMEHeader{
		"Content-Disposition": []string{`form-data; name="file"; filename="testfile1.txt"`}}
	// Test ProcessFiles with an invalid file header (file that does not exist)
	invalidFileHeader := &multipart.FileHeader{
		Filename: "nonexistent.txt",
		Size:     0,
		Header:   h,
	}
	files = []*multipart.FileHeader{invalidFileHeader}
	_, err = ProcessFiles(files, publicKey)
	assert.Error(t, err)
}
