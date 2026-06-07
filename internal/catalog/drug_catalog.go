package catalog
import (
	"stockist-parser/internal/models"
)
// BuildDrugCatalog creates a unique drug catalog
// with generated medicine IDs.
func BuildDrugCatalog(items []models.InventoryItem) []models.Drug {
	var drugs []models.Drug

	for i, item := range items {
		drug := models.Drug{
			MedicineID:     i + 1,
			DrugName:       item.ItemName,
			NormalizedName: item.ItemName,
		}
		drugs = append(drugs, drug)
	}

	return drugs
}