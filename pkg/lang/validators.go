package lang

import (
	"strconv"
	"strings"
)

func IsValidJavaIdentifier(fieldName string) bool {
	return !strings.Contains(fieldName, " ")
}


func DetermineJavaType(rawValue string) JavaType {
	// TODO Refactor
	intResult, err := strconv.ParseInt(rawValue, 0, 64)
	if err == nil {
		if intResult > 2147483647 || intResult < -2147483648 {
			return JavaLong{
				Value: intResult,
			}
		} else {
			return JavaInt{Value: int32(intResult)}
		}
	}
	floatResult, err := strconv.ParseFloat(rawValue, 64)
	if err == nil {
		return JavaDouble{
			Value: floatResult,
		}
	}
	boolResult, err := strconv.ParseBool(rawValue)
	if err == nil {
		return JavaBool{
			Value: boolResult,
		}
	}
	return JavaString{Value: rawValue}
}