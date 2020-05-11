package main

//直接cd test 然后 go test或go test -v
import "testing"

func TestAdd(t *testing.T) {
	r := add(2, 4)
	if r != 6 {
		t.Fatalf("add(2,4) error,expect:%d,actual:%d",6,r)
	}
	t.Logf("test add succ")
}
