package lang

import "fmt"

type Enum struct {
	Id   string
	Repr string
}

func getSeperator(idx, length int) string {
	if idx == length {
		return ";"
	}
	return ","
}

func ConstructEnum(enums []Enum) string {
	result := ""
	for idx, enum := range enums {
		result += fmt.Sprintf("%s(%s)%s\n", enum.Id, enum.Repr, getSeperator(idx, len(enums)-1))
	}
	return result
}
