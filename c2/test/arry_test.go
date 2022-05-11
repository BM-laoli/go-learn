package test

import "testing"

func TestNeweggArray(t *testing.T) {
	// 我们具备如下的声明方式
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3, 4, 5}
	arr1[1] = 0

	t.Log(arr1[1], arr1[2])
	t.Log(arr2)

}

func TestNeweggArrayOp(t *testing.T) {
	// 我们具备如下的 截取和便利取值 赋值 操作
	arr := [...]int{1, 3, 4, 5}

	// 类似js的 forEach 遍历 ，如果你希望省略 某值 ，只需要把它变成 _ 就可以了
	// for idx, e := range arr {
	for _, e := range arr {
		t.Log(e)
	}

	// 快速 的截取
	// a[开始的索引（包含）, 结束索引（不包含）]
	a := [...]int{1, 2, 3, 4, 5}

	t.Log(a[1:2]) // 2
	t.Log(a[1:3]) // 2,3
	t.Log(a[1:])  // 2,3,4,5
	t.Log(a[1:])  // 2,3,4,5
	t.Log(a[:3])  // 1,2,3
}
