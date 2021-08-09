package lang

import (
	"fmt"
	"strconv"
)

type JavaType interface {
	GetTypeName() string
	AsStringValue() string
}

type JavaLong struct {
	Value int
}

func (jl JavaLong) GetTypeName() string {
	return "Long"
}

func (jl JavaLong) AsStringValue() string {
	return fmt.Sprintf("%sL", strconv.Itoa(jl.Value))
}

type JavaString struct {
	Value string
}

func (js JavaString) GetTypeName() string {
	return "String"
}

func (js JavaString) AsStringValue() string {
	return fmt.Sprintf(`"%s"`, js.Value)
}
