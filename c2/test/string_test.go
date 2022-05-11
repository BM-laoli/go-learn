package test

import (
	"strings"
	"testing"
)

func TestStringS(t *testing.T) {
	var s string

	t.Log(s) // 初始化默认值 = 0
	s = "hello"
	t.Log(len(s)) // 对于英语来说 一个字母算一个
	// s[1] = "3" 不能这样写string的本质是一个不可变的 byte slice
	s = "\xE4\xB8\xA5" // 存任何 二进制的数据 这个二进制编码代表 严
	// s = "\xE4\xB8\xA5\xFF" // 存任何 二进制的数据 这个是一个乱码
	t.Log(s)
	t.Log(len((s)))

	s = "中"
	t.Log(len((s))) // 获取的是byte数

	c := []rune(s)               // rune slice 表示的是unicode slice
	t.Logf("中 unicode %x", c[0]) // 使用16 进位制 去获取 0位置上的 unicode编码
	t.Logf("中 UTF %x", s)        // 使用16 进位制 去获取 s
}

func TestStringFN(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	t.Log(parts)

	for _, part := range parts {
		t.Log(part) // 遍历 string
	}
	t.Log(strings.Join(parts, "-")) // 组合
}
