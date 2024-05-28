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
    - `publicKey`: PGP public key for encryption (-----BEGIN PGP PUBLIC KEY BLOCK-----)
    - `bodyText`: Body text of the email
    - `files`: Files to be encrypted and sent

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
  -F "publicKey=-----BEGIN PGP PUBLIC KEY BLOCK-----" \
  -F "bodyText=Please find the encrypted files attached." \
  -F "files=@path/to/file1.txt" \
  -F "files=@path/to/file2.txt"
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## Contact

For any questions or issues, please open an issue on GitHub.
