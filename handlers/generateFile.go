package handlers

import (
	"gin-microservice/pdf"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Payload struct for incoming POST requests
type Payload struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Handler for generating PDF
func GeneratePDFHandler(c *gin.Context) {
	var payload Payload
	// get query params from request
	name := c.Query("name")
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Generate PDF
	pdfData, err := pdf.GeneratePDF(payload.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate PDF"})
		return // Return to stop execution
	}

	// Check if name is empty
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}

	if name == "save" {
		// Save PDF to disk
		filePath, err := pdf.SavePDF(pdfData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save PDF"})
			return
		}
		// Send file path as response
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, gin.H{"pdf_path": filePath})
		return
	} else if name == "download" {
		// Send PDF as response
		c.Header("Content-Type", "application/pdf")
		c.Header("Content-Disposition", "attachment; filename=hello.pdf")
		c.Data(http.StatusOK, "application/pdf", pdfData)
		return
	} else {
		// Send BAD request response
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameter 'name' must be 'save' or 'download'"})
		return
	}

}
