package test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)

	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPter := &a // 取地址符号  获取a的地址 引用（一般来说地址是16进制的）
	t.Log(a, aPter)
	t.Logf("%T %T", a, aPter)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")

	//	判断空srtring 就是获取它的lengyh
	t.Log(len(s))
}

func TestMapFunC(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }

	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T) {
	set := map[int]bool{}
	set[1] = true

	// hash 可以这样实现
	n := 1
	if set[n] {
		t.Log("存在")
	} else {
		t.Log("不存在")
	}

	// delete可以这样实现
	delete(set, n)
	if set[n] {
		t.Log("存在")
	} else {
		t.Log("不存在")
	}
}
