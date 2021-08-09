package lang

import (
	"fmt"
	"testing"
)

func TestConstructEnum(t *testing.T) {
	result := ConstructEnum([]Enum{
		{Id: "US", Repr: `"United States", 101010L`},
		{Id: "RU", Repr: `"Russia", 1000L`},
	})
	fmt.Println(result)
}
