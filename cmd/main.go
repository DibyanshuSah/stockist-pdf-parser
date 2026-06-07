package main

import (
	"fmt"

	"stockist-parser/internal/catalog"
	"stockist-parser/internal/normalizer"
	"stockist-parser/internal/parser"
	"stockist-parser/internal/pdf"
	"stockist-parser/internal/utils"
)

func main() {

	pdfPath := "uploads/demo.PDF"

	// Step 1: Validate PDF
	err := utils.ValidatePDF(pdfPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Step 2: Extract raw text from PDF
	rawText, err := pdf.ExtractText(pdfPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Step 3: Save extracted raw text
	err = utils.SaveText("output/raw.txt", rawText)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Step 4: Parse medicine names
	medicines := parser.ParseMedicines(rawText)

	err = utils.SaveJSON("output/medicines.json", medicines)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Step 5: Remove exact duplicates
	uniqueMedicines := normalizer.BuildUniqueCatalog(medicines)

	err = utils.SaveJSON("output/unique_medicines.json", uniqueMedicines)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Step 6: Generate drug catalog
	drugCatalog := catalog.BuildDrugCatalog(uniqueMedicines)

	err = utils.SaveJSON("output/drug_catalog.json", drugCatalog)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Summary
	fmt.Println()
	fmt.Println("Total Medicines :", len(medicines))
	fmt.Println("Unique Medicines:", len(uniqueMedicines))
	fmt.Println("Drug Catalog Size:", len(drugCatalog))
	fmt.Println()

	limit := 20

	if len(drugCatalog) < limit {
		limit = len(drugCatalog)
	}

	for i := 0; i < limit; i++ {
		fmt.Printf("%d. %s\n", i+1, drugCatalog[i].DrugName)
	}

	fmt.Println()
	fmt.Println("raw.txt saved successfully")
	fmt.Println("medicines.json saved successfully")
	fmt.Println("unique_medicines.json saved successfully")
	fmt.Println("drug_catalog.json saved successfully")
}