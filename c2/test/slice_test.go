package test

import "testing"

// 实际上 slice 是一种结构 包含了 三个重要的 属性 1，地址，2. len  3, cap 容量
func TestSliceNew(t *testing.T) {
	t.Logf("TestSliceNewegg")
	// 我们看看如何 new 一个简单的 slice , 记住它是一个 可变长的 array
	var s0 []int
	t.Log(len(s0), cap(s0))

	// 追加一个如何追加，为什么需要重新 对s0 复制一遍？
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5) // 使用 make 方法创建slice参数 1 是类型 ，参数2是len  ，参数3是容量
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2]) // 注意⚠️  不能访问 len  ，外的 值 会报错！，例如==> s2[4]

}

// 我们看看 slice 的 容量的cap 的增长规律
func TestSliceNewLeng(t *testing.T) {
	s := []int{}

	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log("--->", len(s), cap(s))
		// 通过 图 我们可以看到 cap 的可变范围 都是 * 2 的增长的
	}
}

func TestSlicShaerMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))

	sumer := year[5:8]
	t.Log(sumer, len(sumer), cap(sumer))

	// 用于 是引用 将会导致 值的同步变更
	sumer[0] = "unknown"
	t.Log(Q2, year)

}
