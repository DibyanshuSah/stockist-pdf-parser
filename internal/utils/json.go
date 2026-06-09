package utils
import (
	"encoding/json"
	"os"
)
// SaveJSON writes the data into a formatted JSON style file.
func SaveJSON(path string, data interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, jsonData, 0644)
}

