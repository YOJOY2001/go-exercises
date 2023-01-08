package exce

func PalinDrome(s []string) []string{
	sl := []string{}
	for i := 0; i < len(s); i++ {
		str:=string(s[i])
		rev := ""
		for j := (len(str)-1); j >=0; j-- {
			rev = rev + string(str[j])
		}
		if rev == s[i] {
			sl = append(sl, s[i])
		}
	}
	return sl
}