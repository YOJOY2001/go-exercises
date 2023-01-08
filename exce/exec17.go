package exce

func MatchLetter(str []string, s string) []string {

	sl:=[]string{}

	for _, l := range str {

		for i := 0; i < len(l); i++ {
			w := string(l[i])
			if w == s {
				sl = append(sl, l)
				break
			}
		}
	}
	return sl

}