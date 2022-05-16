package test

import (
	"fmt"
	"testing"
)

func Sumx(arg ...int) int {
	res := 0
	for _, i := range arg {
		res = i + res
	}
	return res
}
func TestMorArg(t *testing.T) {
	t.Log("===>", Sumx(1, 1, 1))
	t.Log("===>", Sumx(1, 2, 3))
}

func print666() {
	fmt.Printf("我是print666")
}
func TestDefer(t *testing.T) {
	defer print666()
	t.Log("我是1 -1 ")
	panic("err")
	t.Log("我是1 -2")

}
