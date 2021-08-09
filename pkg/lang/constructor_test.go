package lang

import (
	"fmt"
	"testing"
)

func TestCreatEnumConstructor(t *testing.T) {

	result := CreatEnumConstructor("ExampleEnum", []EnumField{
		{
			Name: "Country",
			Type: JavaString{},
		},
		{
			Name: "GDP",
			Type: JavaLong{},
		},
	})
	fmt.Println(result)
}

func TestCreateClassVariables(t *testing.T) {
	result := CreateClassVariables([]EnumField{
		{
			Name: "Country",
			Type: JavaString{},
		},
		{
			Name: "GDP",
			Type: JavaLong{},
		},
	})
	fmt.Println(result)
}
