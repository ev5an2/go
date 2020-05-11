package main

import "fmt"

func charu(a []int){
	for i:=1;i<len(a);i++{
		for j:=i;j>0;j--{
			if a[j]>a[j-1] {
				break
			}
			a[j],a[j-1] = a[j-1],a[j]
		}
	}
}

func main(){
	b:=[...]int{3,5,1,8,10}
	charu(b[:])
	fmt.Println(b)
}
