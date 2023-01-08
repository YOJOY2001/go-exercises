package exce

func AscendingOrder(n []int)bool{
	for i:=1;i<len(n);i++{
		if n[i]<n[i-1]{
			return false
		}
	}
	return true
}