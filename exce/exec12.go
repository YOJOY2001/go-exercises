package exce

func Indexing(n []int,j int)int{
	for i,v:=range n{
		if v==j{
			return i
		}
	}
	return 0
}