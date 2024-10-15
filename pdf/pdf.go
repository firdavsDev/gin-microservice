package pdf

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/jung-kurt/gofpdf"
)

// GeneratePDF creates a simple PDF using the provided name.
func GeneratePDF(name string) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, "+name+"!")

	// Create a buffer to hold the PDF output
	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// randomly generated file name using date and time
func GenerateFileName() string {
	// Get the current date and time
	t := time.Now()
	// add the current date and time to the file name (file_2024_12_31_24_00_00) = file_YYYY_MM_DD_HH_MM_SS
	fileName := fmt.Sprintf("file_%d_%02d_%02d_%02d_%02d_%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	return fileName
}

// SavePDF saves the provided PDF data to disk and returns the file path.
func SavePDF(pdfData []byte) (string, error) {
	// Define the directory to save the PDF files
	dir := "media/pdfs" // Adjust this path as needed

	// Create the directory if it doesn't exist
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", fmt.Errorf("could not create directory: %v", err)
	}
	newFileName := GenerateFileName() + ".pdf"
	// Create the full file path
	filePath := filepath.Join(dir, newFileName)

	// Write the PDF data to a file
	err := os.WriteFile(filePath, pdfData, 0644)
	if err != nil {
		return "", fmt.Errorf("could not write PDF file: %v", err)
	}

	return filePath, nil
}

// GetPDFs returns a list of all PDF files in the media/pdfs directory.
func GetPDFs() ([]string, error) {
	dir := "media/pdfs" // Adjust this path as needed

	// Open the directory
	d, err := os.Open(dir)
	if err != nil {
		return nil, fmt.Errorf("could not open directory: %v", err)
	}
	defer d.Close()

	// Read the directory contents
	files, err := d.Readdir(-1)
	if err != nil {
		return nil, fmt.Errorf("could not read directory: %v", err)
	}

	// Filter out non-PDF files
	var pdfs []string
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".pdf" {
			pdfs = append(pdfs, file.Name())
		}
	}

	return pdfs, nil
}

// GetPDF returns the PDF file data for the specified file name.
func GetPDF(fileName string) ([]byte, error) {
	dir := "media/pdfs" // Adjust this path as needed
	new_fileName := fileName + ".pdf"
	filePath := filepath.Join(dir, new_fileName)

	// Read the file data
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %v", err)
	}

	return data, nil
}
