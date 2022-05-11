package test

import (
	"testing"
)

func TestIniitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1)) // f 可以组合模板字符串

	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))

	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))
}

func TestAceessNotEXistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0

	// 注意值不存在的时候是一个 0 ，key存在value = 0 的时候也是o 这个时候我们需要这样判断
	if _, ok := m1[3]; ok {
		t.Logf("key 3's value is %d", m1[3])
	} else {
		t.Log("key 3 is not existing")
	}
}

func TestFoEatch(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for key, value := range m1 {
		t.Log(key, value)
	}
}
