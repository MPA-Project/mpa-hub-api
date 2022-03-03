package utils

func EnumContains(enum []string, value string) bool {
	for _, v := range enum {
		if v == value {
			return true
		}
	}
	return false
}
