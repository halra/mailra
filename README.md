# PGP Email Service

This is a Go-based service for uploading files, encrypting them using PGP (Pretty Good Privacy), and sending them via email. It uses the `gin` web framework for handling HTTP requests and the `gomail` package for sending emails. The PGP encryption is handled by the `gopenpgp` package.

## Features

- Upload multiple files via a POST request.
- Encrypt files using a provided PGP public key.
- Send encrypted files via email, either inline or as MIME attachments.

## Requirements

- Go 1.16 or later
- `gin-gonic/gin` for the web framework
- `gopkg.in/gomail.v2` for sending emails
- `ProtonMail/gopenpgp` for PGP encryption

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/pgp-email-service.git
    cd pgp-email-service
    ```

2. Install the dependencies:
    ```sh
    go mod tidy
    ```

## Configuration

Ensure you have the necessary environment variables set up for your SMTP server configuration. You can also update the default values directly in the code if needed.

## Usage

1. Run the server:
    ```sh
    go run main.go
    ```

2. Make a POST request to the `/api/v1/mail/pgp` endpoint to upload files, encrypt them, and send via email. The request should include the following form fields:

    - `from`: Sender's email address
    - `to`: Recipient's email address
    - `subject`: Email subject
    - `smtPServer`: SMTP server address
    - `smtpPort`: SMTP server port
    - `smtpPassword`: SMTP server password
    - `smtpUser`: SMTP server user
    - `method`: Email sending method (`inline` or `mime`)
    - `publicKey`: PGP public key for encryption
    - `bodyText`: Body text of the email
    - Files: Files to be encrypted and sent

### Example cURL Command

```sh
curl -X POST http://localhost:8080/api/v1/mail/pgp \
  -F "from=your-email@example.com" \
  -F "to=recipient@example.com" \
  -F "subject=Encrypted Files" \
  -F "smtPServer=smtp.example.com" \
  -F "smtpPort=587" \
  -F "smtpPassword=yourpassword" \
  -F "smtpUser=youruser" \
  -F "method=mime" \
  -F "publicKey=@path/to/publicKey.asc" \
  -F "bodyText=Please find the encrypted files attached." \
  -F "files=@path/to/file1.txt" \
  -F "files=@path/to/file2.txt"
```

## Functions

### `pgpHandler`

Handles the main logic for receiving the form data, processing the files, encrypting them, and sending the email.

### `processFiles`

Processes multiple files, encrypting each one using the provided PGP public key.

### `processFile`

Handles the encryption of a single file.

### `encryptPGP`

Encrypts data using the provided PGP public key.

### `sendEmailMIME`

Sends an email with encrypted attachments as MIME.

### `sendEmailInline`

Sends an email with encrypted attachments inline.

### `sendEmail`

Handles the actual sending of the email.

### `createTempFilePath`

Creates a temporary file path for storing encrypted attachments.

### `logAndRespond`

Logs a message and sends an HTTP response.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## Contact

For any questions or issues, please open an issue on GitHub.

```

