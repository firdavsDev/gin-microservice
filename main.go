package main

import (
	"gin-microservice/handlers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set Gin mode from environment variable (optional)
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = "debug" // Default to debug if not set
	}
	gin.SetMode(ginMode)

	r := gin.Default()

	// Register routes

	// Group routes under /api/v1/
	api := r.Group("/api/v1")
	{
		api.POST("/generate-pdf", handlers.GeneratePDFHandler)
		api.GET("/get-pdfs", handlers.GetPDFsHandler)
		api.GET("/get-pdf", handlers.GetPDFHandler)
	}

	// Start server
	log.Println("Starting Gin microservice on :8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed:", err)
	}
}
