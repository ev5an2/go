package main

import "fmt"

func qsort(a []int, left, right int) {
	if left >= right {
		return
	}

	val := a[left]
	k := left
	for i := left + 1; i <= right; i++ {
		if a[i] < val {
			a[k] = a[i]
			a[i] = a[k+1]
			k++
		}
	}
	//3, 2, 1, 8, 10, 2
	//2, 2, 1, 8, 10, 2
	//2, 1, 1, 8, 10, 2
	//2, 1, 2, 8, 10, 2
	//2, 1, 2, 8, 10, 8
	//2, 1, 2, 3, 8, 10
	a[k] = val
	qsort(a, left, k-1)
	qsort(a, k+1, right)

}
func main() {
	b := [...]int{3, 2, 1, 8, 10}
	qsort(b[:], 0, len(b)-1)
	fmt.Println(b)

}
