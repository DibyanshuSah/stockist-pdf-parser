package normalizer

import (
	"regexp"
	"strings"
)
var (
	multiSpaceRegex = regexp.MustCompile(`\s+`)
	specialCharRegex = regexp.MustCompile(`[-_/]+`)
)
// CleanText performs basic cleanup on extracted text.
func CleanText(text string) string {
	text = strings.TrimSpace(text)
	text = strings.ToUpper(text)
	text = specialCharRegex.ReplaceAllString(text, "")
	text = multiSpaceRegex.ReplaceAllString(text, " ")

	return text
}
// NormalizeDosageForm converts common dosage variations
// into a single standard value.
func NormalizeDosageForm(form string) string {
	form = CleanText(form)

	switch form {

	case "TABLET", "TABLETS", "TABS":
		return "TAB"

	case "CAPSULE", "CAPSULES", "CAP":
		return "CAP"

	case "SYRUP", "SYP":
		return "SYP"

	case "INJECTION", "INJ":
		return "INJ"

	case "DROPS", "DROP":
		return "DROP"

	case "CREAM":
		return "CREAM"

	case "OINTMENT":
		return "OINTMENT"

	default:
		return form
	}
}
// NormalizeMedicineName prepares medicine names
// for matching and searching.
func NormalizeMedicineName(name string) string {
	name = CleanText(name)

	return name
}

