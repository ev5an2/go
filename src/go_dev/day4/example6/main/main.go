package main
//数组及初始化
import "fmt"


func fab(n int){
	var a []int
	a = make([]int,n)
	a[0]=1
	a[1]=1
	for i:=2;i<n;i++{
		a[i] = a[i-1]+a[i-2]

	}
	for _,v:=range a{
		fmt.Println(v)
	}
}

func testArry(){
	var a [2][5]int = [...][5]int{{2,3,4,5,6},{1,3,5,7,9}}
	for row,v:=range a{
		for col,v1:=range v{
			fmt.Printf("(%d,%d)-%d",row,col,v1)
		}
	}
}
func main() {
	fab(10)
	testArry()
	var a [3]int = [3]int{1,2,3}
	var b = [3]int{1,2,3}
	var c = [...]int{1,2,3}
	var d = [3]string{1:"d",2:"ddd"}
	var e = [2][2]int{{1,2},{3,4}}
	var f [2][3]int
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
