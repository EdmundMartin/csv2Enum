package main

import (
	"flag"
	"fmt"
	"github.com/EdmundMartin/csv2Enum/pkg/codegen"
)

func main() {
	var enumName string
	var packageName string
	flag.StringVar(&enumName, "enumName", "", "Name of enum")
	flag.StringVar(&packageName, "package", "", "name of package")

	result, _ := codegen.ReadCsvFile("example.csv")
	fmt.Println(result.JavaTypes)

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
