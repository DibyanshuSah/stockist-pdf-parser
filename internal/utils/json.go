package utils
import (
	"encoding/json"
	"os"
)
// SaveJSON writes data into a formatted JSON file.
func SaveJSON(path string, data interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, jsonData, 0644)
}