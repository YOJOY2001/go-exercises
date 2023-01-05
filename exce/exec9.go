package exce

func Average(n[]int)int{
	avg:=0
	i:=0
	for _,v:=range n{
		avg+=v
		i++
	}
	avg=avg/i
	return avg
}