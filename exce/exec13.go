package exce

func Repnum(n []int,j int)int{
	i:=0
	for _,v:=range n{
		if v==j{
			i++
		}
	}
	return i
}