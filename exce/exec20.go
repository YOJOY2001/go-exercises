package exce


func Prime(n []int) []int {
	res:=[]int{}
	for i:=0;i<len(n);i++{
		flag:=0
		if n[i]==0 || n[i]==1{
			flag=1
		}
		for j:=2;j<=n[i]/2;j++{
			if n[i]%j==0{
				flag=1
				break
			}
		}
		if flag == 0{
			res=append(res,n[i])
		}
	}
	return res
}