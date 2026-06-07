package models

// Drug represents a unique medicine
// stored in the master drug catalog.
type Drug struct {
	MedicineID    int    `json:"medicine_id"`
	DrugName      string `json:"drug_name"`
	NormalizedName string `json:"normalized_name"`
}