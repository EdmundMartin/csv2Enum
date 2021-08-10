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
	Value int64
}

func (jl JavaLong) GetTypeName() string {
	return "Long"
}

func (jl JavaLong) AsStringValue() string {
	return fmt.Sprintf("%sL", strconv.Itoa(int(jl.Value)))
}

func (jl JavaLong) String() string {
	return fmt.Sprintf("Long{%d}", jl.Value)
}

type JavaString struct {
	Value string
}

func (js JavaString) GetTypeName() string {
	return "String"
}

func (js JavaString) String() string {
	return fmt.Sprintf("String{%s}", js.Value)
}

func (js JavaString) AsStringValue() string {
	return fmt.Sprintf(`"%s"`, js.Value)
}

type JavaInt struct {
	Value int32
}

func (ji JavaInt) GetTypeName() string {
	return "Integer"
}

func (ji JavaInt) AsStringValue() string {
	return strconv.Itoa(int(ji.Value))
}

func (ji JavaInt) String() string {
	return fmt.Sprintf("Integer{%d}", ji.Value)
}

type JavaDouble struct {
	Value float64
}

func (jd JavaDouble) GetTypeName() string {
	return "Double"
}

func (jd JavaDouble) AsStringValue() string {
	return strconv.FormatFloat(jd.Value, 'f', -1, 64)
}

func (jd JavaDouble) String() string {
	return fmt.Sprintf("Double{%f}", jd.Value)
}

type JavaBool struct {
	Value bool
}

func (jb JavaBool) GetTypeName() string {
	return "Boolean"
}

func (jb JavaBool) AsStringValue() string {
	return fmt.Sprintf("%s", jb.Value)
}

func (jb JavaBool) String() string {
	return fmt.Sprintf("%s", jb.Value)
}