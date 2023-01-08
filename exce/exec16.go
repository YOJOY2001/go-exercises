package exce

func CheckOccur(str []string, s string) int {

	for k, v := range str {
		if v == s {
			return k
		}
	}
	return 0
}