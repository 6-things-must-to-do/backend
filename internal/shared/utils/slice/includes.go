package sliceUtil

func Includes(slice []interface{}, target interface{}) bool {
	for _, val := range slice {
		if target == val {
			return true
		}
	}
	return false
}
