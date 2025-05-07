package utils

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

// SaveFile saves the uploaded file to the specified directory and returns the file path
func SaveFile(uploadDir string, fileHeader *multipart.FileHeader) (string, error) {
	err := os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("failed to create upload dir: %w", err)
	}

	src, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer src.Close()

	ext := filepath.Ext(fileHeader.Filename)
	uniqueFilename := fmt.Sprintf("%d-%d%s", time.Now().Unix(), time.Now().UnixNano()%1e6, ext)
	savePath := filepath.Join(uploadDir, uniqueFilename)

	fileBytes, err := ioutil.ReadAll(src)
	if err != nil {
		return "", fmt.Errorf("failed to read uploaded file: %w", err)
	}

	err = ioutil.WriteFile(savePath, fileBytes, 0644)
	if err != nil {
		return "", fmt.Errorf("failed to save uploaded file: %w", err)
	}

	// Return relative path
	return savePath, nil
}

// DeleteFile deletes the given file if it exists
func DeleteFile(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil // File doesn't exist — no error
	}
	return os.Remove(path)
}
