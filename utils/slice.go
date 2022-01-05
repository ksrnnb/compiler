package utils

func IsInSlice(sl []string, needle string) bool {
	for _, v := range sl {
		if v == needle {
			return true
		}
	}

	return false
}
