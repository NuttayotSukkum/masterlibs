package stringUtils

func IsBlank(val string) bool {
	strLen := len(val)
	if strLen == 0 {
		return true
	} else {
		for i := 0; i < strLen; i++ {
			if string(val[i]) == "" {
				return false
			}
		}
		return true
	}
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
