package exce

func SecLargest(n[]int)int{
	large1:=0
	large2:=0
	for _,v:=range n{
		if large1<v{
			large2=large1
			large1=v
		}
		if v>large2 && v<large1{
			large2=v
		}
	}
	return large2
}