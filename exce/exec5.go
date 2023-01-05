package exce

func Largest(n[]int)int{
	large:=0
	for _,v:=range n{
		if large<v{
			large=v
		}
	}
	return large
}