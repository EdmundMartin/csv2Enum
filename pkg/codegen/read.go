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
	return CsvData{headers: lines[0], rows: lines[1:]}, nil
}
