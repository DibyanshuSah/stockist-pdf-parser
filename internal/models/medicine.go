package models
// InventoryItem represents one row extracted from a PDF.
// Different stockists may use different column names.
// Example:
// Medicine Name
// Product Name
// Item Description
// Drug Name     (all of them will eventually map into these fields.)
type InventoryItem struct {
	ItemName string `json:"item_name"`
	ItemCode string `json:"item_code"`
	GenericName string `json:"generic_name"`
	Composition string `json:"composition"`
	Manufacturer string `json:"manufacturer"`
	Strength string `json:"strength"`
	DosageForm string `json:"dosage_form"`
	Category string `json:"category"`
	HSNCode string `json:"hsn_code"`
	GSTPercent string `json:"gst_percent"`
	Quantity string `json:"quantity"`
	Unit string `json:"unit"`
	BatchNo string `json:"batch_no"`
	ExpiryDate string `json:"expiry_date"`
	MRP string `json:"mrp"`
	RawText string `json:"raw_text"`
}

