package handlers

/*
GetPDFHandler function is a handler for getting PDF files by file name in media/pdf folder via GET request to /get-pdf endpoint.
GetPDFsHandler function is a handler for getting all PDF files in media/pdf folder names list via GET request to /get-pdfs endpoint.
*/

import (
	"gin-microservice/pdf"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPDFsHandler function is a handler for getting all PDF files in media/pdf folder names list via GET request to /get-pdfs endpoint.
func GetPDFsHandler(c *gin.Context) {
	// Get all PDF files in media/pdf folder
	files, err := pdf.GetPDFs()
	if err != nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get PDFs"})
		return
	}

	// Send files list as response
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"pdfs": files})
	return
}

// GetPDFHandler function is a handler for getting PDF files by file name in media/pdf folder via GET request to /get-pdf endpoint.
func GetPDFHandler(c *gin.Context) {
	// Get file name from query parameter
	fileName := c.Query("file")
	if fileName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File name is required"})
		return
	}

	// Get file path
	data, err := pdf.GetPDF(fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get PDF"})
		return

	}
	// Send file data as response
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Data(http.StatusOK, "application/pdf", data)
	return

}
