package util

// StringSliceContains receives a slice of string and checks if contains str
func StringSliceContains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}

	return false
}
