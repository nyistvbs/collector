package helper

func InStringArray(arr []string, str string) bool {
	for _, v := range arr {
		if str == v {
			return true
		}
	}
	return false
}
