package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/halra/mailra/domain"

	"github.com/halra/mailra/utils"
)

func PgpHandler(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to parse form: %s", err.Error()))
		return
	}
	files := form.File["files"]

	mr := domain.MailRequest{
		From:         c.PostForm("from"),
		To:           c.PostForm("to"),
		Subject:      c.PostForm("subject"),
		SmtPServer:   c.PostForm("smtPServer"),
		SmtpPort:     c.PostForm("smtpPort"),
		SmtpPassword: c.PostForm("smtpPassword"),
		SmtpUser:     c.PostForm("smtpUser"),
		Method:       c.PostForm("method"),
		BodyText:     c.PostForm("bodyText"),
		PublicKey:    c.PostForm("publicKey"),
	}

	attachments, err := utils.ProcessFiles(files, mr.PublicKey)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to process file: %s", err.Error()))
		return
	}

	if mr.Method == "inline" {
		log.Println("Sending email with inline attachments")
		err = utils.SendEmailInline(attachments, mr)
	} else {
		log.Println("Sending email with MIME attachments")
		err = utils.SendEmailMIME(attachments, mr)
	}

	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to send email: %s", err.Error()))
		return
	}

	c.String(http.StatusOK, "Files uploaded and sent successfully.")
}
