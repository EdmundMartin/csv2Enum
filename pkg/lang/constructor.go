package lang

import (
	"fmt"
	"strings"
)

const constructor = "%s(%s){"

type EnumField struct {
	Name string
	Type JavaType
}

func generateConstructorArgs(fields []EnumField) string {
	fieldRepr := ""
	for _, val := range fields {
		fieldRepr += fmt.Sprintf("final %s %s,", val.Type.GetTypeName(), val.Name)
	}
	fieldRepr = strings.TrimRight(fieldRepr, ",")
	return fieldRepr
}

func addConstructorSetters(fields []EnumField) string {
	result := "\n"
	for _, val := range fields {
		result += fmt.Sprintf("this.%s = %s;\n", val.Name, val.Name)
	}
	result += "}"
	return result
}

func CreatEnumConstructor(name string, fields []EnumField) string {
	enum := fmt.Sprintf(constructor, name, generateConstructorArgs(fields))
	return enum + addConstructorSetters(fields)
}

func CreateClassVariables(fields []EnumField) string {
	result := ""
	for _, field := range fields {
		result += fmt.Sprintf("private %s %s;\n", field.Type.GetTypeName(), field.Name)
	}
	return result
}
