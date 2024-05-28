package main

import (
	"github.com/gin-gonic/gin"

	"github.com/halra/mailra/controller"
)

func main() {
	router := gin.Default()
	router.POST("/api/v1/mail/pgp", controller.PgpHandler)
	router.Run(":8080")
}
