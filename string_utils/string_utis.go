package string_utils

// 检查slice 是否包含当前字符串
func Contains(slice []string, target string) bool {
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
