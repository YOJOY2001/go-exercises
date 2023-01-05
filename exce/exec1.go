package exce

func Filtereven(slice []int) []int {
	sl := []int{}
	for _, v := range slice {
		if v%2 == 0 {
			sl = append(sl, v)
		}
	}
	return sl
}
