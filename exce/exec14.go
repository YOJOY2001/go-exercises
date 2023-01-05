package exce

func Reverse(n[]int)[]int{
	sl:=[]int{}
	for i:=len(n);i>0;i--{
		sl=append(sl,n[i-1])
	}
	return sl
}