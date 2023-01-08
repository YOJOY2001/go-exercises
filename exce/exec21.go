package exce

func AlternateNum(n []int) []int {
	sl1:= []int{}
	sl2:= []int{}
	sl3:= []int{}
	for i, v := range n {
		if i%2 == 0 {
			sl1 = append(sl1, v)
		} else {
			sl2 = append(sl2, v)
		}
	}
	sl3 = append(sl1, sl2...)

	return sl3
}