package exce

func Check(s []string,l string)bool{
	for _,v:=range s{
		if l==v{
			return true
		}
	}
	return false
}