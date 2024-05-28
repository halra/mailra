package domain

type EncryptedAttachment struct {
	Filename string
	Content  []byte
}

type MailRequest struct {
	From         string
	To           string
	Subject      string
	SmtPServer   string
	SmtpPort     string
	SmtpPassword string
	SmtpUser     string
	Method       string
	PublicKey    string
	BodyText     string
}
