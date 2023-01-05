package exce

func Freq(s string)map[string]int{
	res:=make(map[string]int)
	for i:=0;i<len(s);i++{
		// v,ok:=res[s[i]]
		// if ok{
		// 	res[s[i]]=v+1
		// }else{
		// 	res[s[i]]=1
		// }
	}
	return res
}