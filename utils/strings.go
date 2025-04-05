package utils

type AllowedString string

func IsStringAllowed(pStrings []AllowedString, cString string) bool {
	if len(pStrings) == 0 {
		return false
	}

	for _, allowedStr := range pStrings {
		if string(allowedStr) == cString {
			return true
		}
	}

	return false
}
