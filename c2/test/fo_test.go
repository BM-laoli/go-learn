package test

import (
	"fmt"
	"testing"
)

func TestBaseFor(t *testing.T) {
	// 简单的for
	for i := 0; i < 5; i++ {
		t.Log(i)
	}

	// while循环
	a := 5
	for a < 5 {
		t.Log(a)
		a--
	}

}

func TestIf(t *testing.T) {
	if v, err := someFunctiono(); err == nil {
		t.Log("没有出错", v)
	} else {
		t.Log("有出错", err)
	}
}

func someFunctiono() (sum int, err error) {
	a := 10
	fmt.Println(a)
	return a, nil
}

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}

func TestSwitchConditionCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("unknown")
		}
	}
}
