package codegen

import (
	"encoding/csv"
	"fmt"
	"github.com/EdmundMartin/csv2Enum/pkg/lang"
	"os"
)

type CsvData struct {
	headers []string
	rows    [][]string
	JavaTypes [][]lang.JavaType
}

func NewCsvData(headers []string, rows [][]string) CsvData {
	return CsvData{
		headers:   headers,
		rows:      rows,
		JavaTypes: [][]lang.JavaType{},
	}
}

func (c CsvData) ValidateHeaders() bool {
	hasName := false
	for _, val := range c.headers {
		if !lang.IsValidJavaIdentifier(val) {
			return false
		}
		if val == "name" {
			hasName = true
		}
	}
	return hasName
}

func (c CsvData) ValidateColumnLengths() bool {
	headerSize := len(c.headers)
	for _, row := range c.rows {
		if len(row) != headerSize {
			return false
		}
	}
	return true
}

func (c *CsvData) CalculateJavaTypes() {
	for _, row := range c.rows {
		var rowTypes []lang.JavaType
		for _, value := range row {
			rowTypes = append(rowTypes, lang.DetermineJavaType(value))
		}
		c.JavaTypes = append(c.JavaTypes, rowTypes)
	}
}


func ReadCsvFile(fileLocation string) (CsvData, error) {
	f, err := os.Open(fileLocation)

	if err != nil {
		return CsvData{}, err
	}
	lines, err := csv.NewReader(f).ReadAll()
	fmt.Println(lines)
	if err != nil {
		return CsvData{}, err
	}
	cData := NewCsvData(lines[0], lines[1:])
	cData.CalculateJavaTypes()
	return cData, nil
}
