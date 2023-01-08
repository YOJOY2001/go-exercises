package exce

import "sort"


func LCM(n []int)int{
	sort.Ints(n)
	i:=n[len(n)-1]
	for{
		flag:=true
		for _,v:=range n{
			if i%v!=0{
				flag=false
				break
			}
		}
		if flag==true{
			break
		}
		i++
	}
	return i
}
