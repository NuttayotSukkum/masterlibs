package stringUtils

import "strings"

func IsBlank(val string) bool {
	return strings.TrimSpace(val) == ""
}

func IsEmpty(val *string) bool {
	return val == nil || len(*val) == 0
}

func IsNotBlank(val string) bool {
	return !IsBlank(val)
}

func IsNotEmpty(val *string) bool {
	return !IsEmpty(val)
}
