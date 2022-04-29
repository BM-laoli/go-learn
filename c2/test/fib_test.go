package test

import (
	"fmt"
	"testing"
)

// 知识点 变量声明
// 测试一个简单的 斐波那契 数列
var a = 0 // 全局的

func TestFbi(t *testing.T) {
	// 声明变量
	// a := 1 局部的
	b := 2

	for i := 0; i < 5; i++ {
		fmt.Println(" ", b)
		temp := a
		a = b
		b = temp + a
	}
}

func TestExchange(t *testing.T) {
	x := 1
	xx := 2

	x, xx = xx, x

	fmt.Println("---> 交换了")
	t.Log(x)
	t.Log(xx)

}
