package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/halra/mailra/controller"
)

func main() {

	// Set the Gin mode based on the environment variable, default to "debug"
	//GIN_MODE=release gin.ReleaseMode

	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)

	router := gin.Default()
	//POST
	router.POST("/api/v1/mail/pgp", controller.PgpHandler)
	//GET
	router.GET("/health", controller.HealthCheckHandler)

	// Define the server address
	addr := ":8080"

	// Print startup information
	printStartupInfo(addr, mode)
	// Start the server
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// printStartupInfo prints server startup information
func printStartupInfo(addr, mode string) {
	fmt.Println("======================================")
	fmt.Println("            Server Startup            ")
	fmt.Println("======================================")
	fmt.Printf("Time: %s\n", time.Now().Format(time.RFC1123))
	fmt.Printf("Version: %s\n", "0.0.1")
	fmt.Printf("Address: %s\n", addr)
	fmt.Printf("Environment: %s\n", mode)
	fmt.Println("======================================")
	fmt.Println("API Endpoints:")
	fmt.Println("  POST /api/v1/mail/pgp")
	fmt.Println("  GET  /health")
	fmt.Println("======================================")
	fmt.Println("Server is starting... Press CTRL+C to stop.")
	fmt.Println("======================================")
}
