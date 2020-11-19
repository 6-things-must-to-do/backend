package sliceUtil

func Includes(slice []string, target string) bool {
	for _, val := range slice {
		if target == val {
			return true
		}
	}
	return false
}
