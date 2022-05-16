package test

import (
	"errors"
	"fmt"
	"testing"
)

//我们通常 这样来处理错误
var Err1 = errors.New("我要大于2")
var Err2 = errors.New("我要小于100")

func getFib(n int) ([]int, error) {
	if n < 2 {
		return nil, Err1
	}
	if n > 100 {
		return nil, Err2
	}
	// 错误优先处理

	fibLit := []int{1, 1}
	for i := 2; i < n; i++ {
		fibLit = append(fibLit, fibLit[i-2]+fibLit[i-1])
	}
	return fibLit, nil
}

func TestErorr(t *testing.T) {
	v, err := getFib(5)
	if err == Err1 {
		t.Log(Err1)
		return
	}

	if err == Err2 {
		t.Log(Err2)
		return
	}

	t.Log(v)

}

func TestErrorCahth(t *testing.T) {
	// fmt.Printf("xxxx")
	// os.Exit(-1)  // 退出不会有任何信息
	// fmt.Printf("1111") // 不会执行

	// defer func() { // defer 不会被error 影响
	// 	fmt.Printf("deferdeferdeferdefer")
	// }()
	// panic(errors.New("error")) // 会有错误信息
	// // 这个函数接受一个 空 interface  你应该理解这是为什么 类似any

	// 下面就是 recover 的用法，但是我们一般不会这样写，因为log 并没有任何意义还会有问题
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("这就是recover的用法", err)
		}
	}()
	panic("我错了")
}
