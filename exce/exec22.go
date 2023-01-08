package exce

func DivisibleByN(n []int, num int) []int {
	res:=[]int{}
	for _, v := range n {
		if v%num == 0 {
			res = append(res, v)
		}
	}
	return res
}