package main

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"
    "os"
    "path/filepath"
    "io"
    "mime/multipart"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func TestUploadExcel(t *testing.T) {
    // Set up Gin router
    router := gin.Default()
    router.POST("/upload", uploadExcel)

    // Create a sample Excel file for testing
    f, err := os.Create("test.xlsx")
    if err != nil {
        t.Fatalf("Failed to create test Excel file: %v", err)
    }
    f.Close()
    defer os.Remove("test.xlsx") // Clean up after the test

    // Prepare the request
    file, err := os.Open("test.xlsx")
    if err != nil {
        t.Fatalf("Failed to open test file: %v", err)
    }
    defer file.Close()

    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)
    part, err := writer.CreateFormFile("file", filepath.Base("test.xlsx"))
    if err != nil {
        t.Fatalf("Failed to create form file: %v", err)
    }
    _, err = io.Copy(part, file)
    if err != nil {
        t.Fatalf("Failed to copy file: %v", err)
    }
    writer.Close()

    req := httptest.NewRequest(http.MethodPost, "/upload", body)
    req.Header.Set("Content-Type", writer.FormDataContentType())

    // Record the response
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    // Assert the response
    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), "Excel file processed successfully")
}
