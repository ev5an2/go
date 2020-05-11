package main

import "fmt"

func choose(a []int){
	for i:=0;i<len(a);i++{
		for j:=i+1;j<len(a);j++{
			if a[i]>a[j] {
				a[i],a[j] = a[j],a[i]

			}
		}
	}
}

func main(){
	b:=[...]int{3,5,1,8,10}
	choose(b[:])
	fmt.Println(b)
}
