package exce

import "strings"

func PrintBackwards(s string) string {
	part := strings.Split(s," ")
	res:=[]string{}
	for i := len(part) - 1; i >= 0; i-- {
		if part[i]==" "{
			continue
		}
		res=append(res, part[i])
	}
	return strings.Join(res, " ")
}

// func reverseString(str string) string {
// 	var ans []string
// 	parts := strings.Split(str, " ")
// 	for i := len(parts) - 1; i >= 0; i-- {
// 		if len(parts[i]) == 0 {
// 			continue
// 		}
// 		ans = append(ans, parts[i])
// 	}
// 	return strings.Join(ans, " ")
// }