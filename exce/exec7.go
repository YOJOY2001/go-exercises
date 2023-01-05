package exce

import "strings"

func Remvow(s string) string {
	for _, v := range []string{"a","e","i","o","u"} {
		s = strings.ReplaceAll(s,v,"")
	}
	return s
}