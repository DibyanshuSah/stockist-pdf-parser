# Stockist PDF Parser

A Go-based PDF parsing system that extracts medicine names from stockist inventory PDFs and generates a unique drug catalog.

## Features

- Extract text from PDF files
- Parse medicine names from stockist reports
- Remove exact duplicate medicines
- Generate a unique drug catalog
- Export data as JSON
- Modular and scalable architecture
- Designed for future PostgreSQL integration

---

## Current Pipeline

```text
PDF
 ↓
Text Extraction
 ↓
Medicine Extraction
 ↓
Medicine JSON
 ↓
Unique Medicine Catalog
 ↓
Drug Catalog
```

---

## Project Structure

```text
stockist-parser/
│
├── cmd/
│   └── main.go
│
├── internal/
│   ├── catalog/
│   │   └── drug_catalog.go
│   │
│   ├── models/
│   │   ├── drug.go
│   │   └── medicine.go
│   │
│   ├── normalizer/
│   │   ├── catalog.go
│   │   └── normalizer.go
│   │
│   ├── parser/
│   │   └── medicine_parser.go
│   │
│   ├── pdf/
│   │   └── extractor.go
│   │
│   └── utils/
│       ├── file.go
│       └── json.go
│
├── uploads/
├── output/
├── go.mod
└── go.sum
```

---

## Tech Stack

- Go
- ledongthuc/pdf
- pdfcpu
- JSON
- PostgreSQL (planned)

---

## Output Files

The parser generates:

```text
raw.txt
medicines.json
unique_medicines.json
drug_catalog.json
```

### Example Drug Catalog Entry

```json
{
  "medicine_id": 1,
  "drug_name": "A TO Z CV CAP",
  "normalized_name": "A TO Z CV CAP"
}
```

---

## Run Locally

Install dependencies:

```bash
go mod tidy
```

Run the project:

```bash
go run cmd/main.go
```

---

## Current Status

### Completed

- PDF text extraction
- Medicine name extraction
- JSON export
- Unique medicine catalog generation
- Drug catalog generation

### Planned

- Improved normalization
- Fuzzy matching
- Multi-PDF support
- PostgreSQL integration
- Stockist ↔ Medicine mapping
- REST APIs
- Frontend dashboard

---

## Future Database Design

### Drugs

```text
medicine_id
drug_name
normalized_name
```

### Stockists

```text
stockist_id
pdf_reference
uploaded_at
```

### Stockist Drug Mapping

```text
stockist_id
medicine_id
```

---
