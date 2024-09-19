package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 10 << 20 // Limit file upload to 10MB
	router.POST("/upload", uploadExcel)

	router.Run(":8080") // Start the server on port 8080
}

// Function to handle file upload
func uploadExcel(c *gin.Context) {
	// Retrieve the uploaded file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File upload failed"})
		return
	}

	// Validate that the file is an Excel file
	if filepath.Ext(file.Filename) != ".xlsx" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file format"})
		return
	}

	// Save the uploaded file
	err = c.SaveUploadedFile(file, file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "File saving failed"})
		return
	}

	// Process the Excel file
	extractDataFromExcel(file.Filename, c)
}

// Function to extract data from the Excel file
func extractDataFromExcel(filePath string, c *gin.Context) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatal("Error opening Excel file:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open Excel file"})
		return
	}
	defer f.Close()

	// Get all sheet names
	sheetNames := f.GetSheetList()
	if len(sheetNames) == 0 {
		log.Fatal("No sheets found in Excel file")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No sheets found in Excel file"})
		return
	}

	// Use the first available sheet
	sheetName := sheetNames[0]

	// Read all rows from the first sheet
	rows, err := f.GetRows(sheetName)
	if err != nil {
		log.Fatal("Error reading rows:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read Excel data"})
		return
	}

	// Process the rows (for now, just logging the rows)
	for _, row := range rows {
		log.Println(row)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Excel file processed successfully"})
}
