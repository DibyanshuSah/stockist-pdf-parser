package pdf
import (
	"fmt"
	"strings"
	pdf "github.com/ledongthuc/pdf"
)
// ExtractText helps in the ectraction of the plain text from all pages of a PDF.
func ExtractText(pdfPath string) (string, error) {
	file, reader, err := pdf.Open(pdfPath)
	if err != nil {
		return "", fmt.Errorf("unable to open pdf: %w", err)
	}
	defer file.Close()

	var textBuilder strings.Builder
	totalPages := reader.NumPage()

	for pageNo := 1; pageNo <= totalPages; pageNo++ {
		page := reader.Page(pageNo)
		if page.V.IsNull() {
			continue
		}
		content, err := page.GetPlainText(nil)
		if err != nil {
			continue
		}
		textBuilder.WriteString(content)
		textBuilder.WriteString("\n")
	}

	return textBuilder.String(), nil
}



