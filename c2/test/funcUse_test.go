package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 测试多返回值
func returnMuintValue() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}
func TestFn(t *testing.T) {
	a, b := returnMuintValue()

	t.Log(a, b)
}

// wrap函数 FP模式
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		startTme := time.Now()
		ret := inner(n)
		fmt.Println("time span ===>", time.Since(startTme).Seconds())
		return ret
	}
}

func showFN(value int) int {
	time.Sleep(time.Second * 1)
	return value
}

func TestWrapFP(t *testing.T) {
	tsSFFP := timeSpent(showFN)
	t.Log(tsSFFP(10))
}
