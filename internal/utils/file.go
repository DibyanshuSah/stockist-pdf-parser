package utils
import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)
// FileExists checks whether a file exists.
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// IsPDF verifies file extension.
func IsPDF(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	return ext == ".pdf"
}

// ValidatePDF performs basic validation.
func ValidatePDF(path string) error {

	if !FileExists(path) {
		return errors.New("file not found")
	}
	if !IsPDF(path) {
		return errors.New("only pdf files are supported")
	}

	return nil
}
// EnsureDirectory creates a folder if needed.
func EnsureDirectory(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}
// SaveText writes extracted text to disk.
func SaveText(path string, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}