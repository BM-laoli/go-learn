package test

import (
	"fmt"
	"testing"
	"time"
)

type Programer interface {
	SayHello() string
}

type GoProgramer struct{}

// 这里要求 darkTyp 签名完全一致  （  	SayHello() string  ）
func (g *GoProgramer) SayHello() string {
	return "fmt.Println(\"Hello!!!!\")"
}

func TestCLinet(t *testing.T) {
	p := new(GoProgramer)

	t.Log(p.SayHello())
}

// 自定义类型
type MyFunction func(op int) int

func timeSpentX(inner MyFunction) MyFunction {
	return func(n int) int {
		startTme := time.Now()
		ret := inner(n)
		fmt.Println("time span ===>", time.Since(startTme).Seconds())
		return ret
	}
}
