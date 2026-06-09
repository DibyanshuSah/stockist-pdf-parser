package parser

import (
	"bufio"
	"regexp"
	"strings"

	"stockist-parser/internal/models"
)
var (
// some common pack formatsle like: 10'S,15S,30ML,100ML,1VIAL,20TAB and 5CAP are 
	packRegex = regexp.MustCompile(`(?i)^(\d+('|’)S|\d+S|\d+ML|\d+GM|\d+KG|\d+TAB|\d+CAP|\d+VIAL|\d+AMP)$`)
	// other things like HSN codes are usually have the long numeric values.
	hsnRegex = regexp.MustCompile(`^\d{4,}$`)
	// for the Tax/composition values
	compositionRegex = regexp.MustCompile(`^\d+(\.\d+)?\+\d+(\.\d+)?$`)
)
// ParseMedicines extracts medicine names from raw pdf text.
func ParseMedicines(rawText string) []models.InventoryItem {
	var medicines []models.InventoryItem
	scanner := bufio.NewScanner(strings.NewReader(rawText))
	startReading := false
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line=="" {
			continue
		}
		upperLine := strings.ToUpper(line)

		if strings.Contains(upperLine, "ITEM DESCRIPTION") {
			startReading=true
			continue
		}
		if !startReading {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}

		medicineName := extractMedicineName(fields)

		if medicineName == "" {
			continue
		}
		medicines = append(medicines, models.InventoryItem{
			ItemName: medicineName,
			RawText:  line,
		})
	}

	return medicines
}

// extractMedicineName keeps the reading until inventory related
// numeric columns starts here.
func extractMedicineName(parts []string) string {
	var nameParts []string
	for _, word := range parts {
		upperWord := strings.ToUpper(word)
		// Pack size starts here.
		if packRegex.MatchString(upperWord) {
			break
		}
		// Composition starts here.
		if compositionRegex.MatchString(upperWord) {
			break
		}

		// HSN starts from here.
		if hsnRegex.MatchString(upperWord) {
			break
		}
		// Common GST values.
		if upperWord == "0" || upperWord == "5" || upperWord == "12" || upperWord == "18" || upperWord == "28"{
			break
		}
		nameParts = append(nameParts, word)
	}
	name := strings.Join(nameParts, " ")
	name = strings.TrimSpace(name)

	return name
}


