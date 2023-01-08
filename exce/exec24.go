package exce

func MinLength(s[]string,l int)[]string{
	sl:=[]string{}
	for _,v:=range s{
		if len(v)>=l{
			sl=append(sl,v)
		}
	}
	return sl
}