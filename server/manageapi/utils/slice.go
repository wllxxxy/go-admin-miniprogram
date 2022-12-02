package utils

func InSlice(val string, arr []string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}

	return false
}
