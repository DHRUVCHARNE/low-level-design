package model

import "fmt"

type DataExporter interface {
	validate(data []string) bool
	Export(data []string)
}
type BaseExporter struct{}
type CSVExporter struct{ b BaseExporter }
type JSONExporter struct{ b BaseExporter }

func NewCSVExporter() *CSVExporter {
	return &CSVExporter{}
}

func NewJSONExporter() *JSONExporter {
	return &JSONExporter{}
}

func (d *BaseExporter) validate(data []string) bool {
	if len(data) == 0 {
		fmt.Println("Export failed: No data to export.")
		return false
	}
	fmt.Println("Validation passed. Exporting N records.")
	return true
}

func (c *JSONExporter) Export(data []string) {
	if !c.b.validate(data) {
		return
	}
	fmt.Printf("JSON: ")
	size := len(data)
	if size == 0 {
		fmt.Println("(empty)")
		return
	}
	fmt.Printf("[")

	for i := 0; i < size-1; i++ {
		fmt.Printf("%s,", data[i])
	}
	fmt.Printf("%s",data[size-1])
	fmt.Printf("]")
}

func (c *CSVExporter) Export(data []string) {
	if !c.b.validate(data) {
		return
	}
	fmt.Printf("CSV: ")
	size := len(data)
	if size == 0 {
		fmt.Println("CSV: (empty)")
		return
	}
	for i := 0; i < size-1; i++ {
		fmt.Printf("%s,", data[i])
	}
	fmt.Println(data[size-1])
}
