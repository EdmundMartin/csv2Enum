package lang

import "strings"

func IsValidJavaIdentifier(fieldName string) bool {
	return !strings.Contains(fieldName, " ")
}
