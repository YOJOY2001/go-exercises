package exce


func PalNum(n []int)[]int{
	res:=[]int{}
	for _, v := range n {
		old:=v
		rem:=0
		rev:=0
		for v != 0 {
			rem = v % 10
			rev = (10 * rev) + rem
			v = v / 10
		}
		if rev==old{
			res=append(res,rev)
		}
	}
	return res
}
