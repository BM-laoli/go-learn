package test

import "testing"

func TestXXX(t *testing.T) {
	t.Log("Testing")
}

// 依次使用 累加 +1
const (
	Monday = 1 + iota
	Tuesday
	Wednesday
)

// 我们使用位运算 来做判断
const (
	Readbale = 1 << iota
	Weitable
	Executable
)

func TestConstAdd(t *testing.T) {
	t.Log("Testing", Monday)
	t.Log("Tuesday", Tuesday)
}

// 使用与运算 它可以 非常建议的表示四种类型
func TestConstAdd2(t *testing.T) {
	testValue := 7 // 0111
	t.Log("Testing", testValue&Readbale)
	t.Log("Testing", testValue&Weitable)
	t.Log("Testing", testValue&Executable)
}
