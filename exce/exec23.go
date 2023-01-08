package exce

func CheckNum(n []int, num int) bool {
	for _, v := range n {
		if v == num {
			return true
		}
	}
	return false
}