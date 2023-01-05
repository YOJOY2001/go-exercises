package exce

func Dupnum(n []int)[]int{
	res:=make(map[int]int)
	sl:=[]int{}
	for _,val:=range n{
		v,ok:=res[val]
		if ok{
			res[val]=v+1
		}else{
			res[val]=1
		}
		if res[val]==1{
			sl=append(sl,val)
		}
	}
	return sl
}