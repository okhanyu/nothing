package util

func ArraysStringContain(strArray []string, target string) bool {
	for _, str := range strArray {
		if str == target {
			return true
		}
	}
	return false
}
