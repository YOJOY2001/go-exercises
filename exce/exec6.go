package exce

func Fivecr(s[]string)[]string{
	sl:=[]string{}
	for _,v:=range s{
		if len(v)>=5{
			sl=append(sl,v)
		}
	}
	return sl
}