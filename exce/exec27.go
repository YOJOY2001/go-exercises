package exce

func ReverseString(str string) string {
	res := ""
	for i := len(str) - 1; i >= 0; i-- {
		res += string(str[i])
	}
	return res
}