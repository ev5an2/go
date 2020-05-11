package main

import "fmt"

func isPrime(a int) bool {
	if (a <= 1) {
		return false
	} else {
		for i := 2; i < a; i++ {
			if (a%i == 0) {
				return false
			}
		}
		return true
	}

}

func main() {
	success:=0
	var n int
	var m int
	fmt.Scanf("%d%d",&n,&m)

	for i := n; i<m;i++{
		if isPrime(i) == true{
			success ++
			fmt.Printf("%d\n",i)
			continue
		}
	}
	fmt.Println("\n",success)
}
