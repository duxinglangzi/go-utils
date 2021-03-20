package int_utils

func Contains(slice []int, target int) bool {
	if len(slice) == 0 {
		return false
	}
	for _, s := range slice {
		if s == target {
			return true
		}
	}
	return false
}
