package exce

func WordCount(s string)map[string]int{
	res:=make(map[string]int)
	for i:=0;i<len(s);i++{
		w:=string(s[i])
		v,ok:=res[w]
		if ok{
			res[w]=v+1
		}else{
			res[w]=1
		}
	}
	return res
}