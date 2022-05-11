package test

import "testing"

func TestOperatorArray(t *testing.T) {
	a := [...]int64{1, 2, 3, 4}
	b := [...]int64{1, 2, 3, 3}
	c := [...]int64{1, 2, 3, 2}
	d := [...]int64{1, 2, 3, 4}

	t.Log(a == b, c == d, a == d)
}

func TestConstAdd3(t *testing.T) {
	const (
		Readbale = 1 << iota
		Weitable
		Executable
	)
	testValue := 7 // 0111
	// 让它变成 不可写
	testValue = testValue &^ Readbale
	// 让它变成 不可读
	// testValue = testValue &^ Weitable
	// 让它变成 不可执行
	testValue = testValue &^ Executable
	t.Log("Testing", testValue&Readbale == Readbale)
	t.Log("Testing", testValue&Weitable == Weitable)
	t.Log("Testing", testValue&Executable == Executable)
}
