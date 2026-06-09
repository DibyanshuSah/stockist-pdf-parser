package normalizer
import (
	"strings"
	"stockist-parser/internal/models"
)
// BuildUniqueCatalog removes the exact duplicates
// after applying the basic normalization.
func BuildUniqueCatalog(items []models.InventoryItem) []models.InventoryItem {

	seen :=make(map[string]bool)
	var uniqueItems []models.InventoryItem
	for _, item := range items {
		name := NormalizeMedicineName(item.ItemName)
		if name==""{
			continue
		}
		if seen[name]{
			continue
		}
		seen[name]=true
		item.ItemName=strings.TrimSpace(name)
		uniqueItems=append(uniqueItems, item)
	}

	return uniqueItems
}



