package main

import (
	"flag"
	"github.com/EdmundMartin/csv2Enum/pkg/codegen"
)

func main() {
	var enumName string
	var packageName string
	flag.StringVar(&enumName, "enumName", "", "Name of enum")
	flag.StringVar(&packageName, "package", "", "name of package")

	codegen.ReadCsvFile("example.csv")

	codegen.GenerateEnum(codegen.EnumInfo{
		Package:  "com.edmundmartin.example",
		EnumName: "Example",
		Entries: []codegen.EnumEntry{
			{
				Id:   "EU",
				Repr: `"ASDSAAS",Region.EU`,
				Done: true,
			},
		},
	})
}
