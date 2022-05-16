package test

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Name string
	Id   string
	Age  int
}

func TestFuncOOP(t *testing.T) {
	e := Employee{"0", "Bod", 20}
	e1 := Employee{Name: "Aoda", Id: "1", Age: 30}

	e2 := new(Employee) // 实际上new 返回的是一个 实例的指针 既 e := &Employee （&取指针操作符在go中） 在c语言中如果你获取到死指针变量 如果你需要获取值 需要 -> 符号但是go中不需要
	e2.Id = "2"
	e2.Name = "Joney"

	t.Log(e)
	t.Log(e1)
	t.Log(e2)

	t.Logf("e is %T", e)   // %T 取类型 test.Employee
	t.Logf("e2 is %T", e2) // *test.Employee

}

// 添加 行为 ，这个方式的话，由于我们前面说了 go 的function  参数都是值拷贝的，你这么写 它会直接在内存中给你搞一份空间 初始化 Employee，如何验证请看下面的代码 x处
// func (e Employee) String() string {
// 	fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))                // 这个 unsafa包自己下去 %x是去16进制的数据，一般来说内存地址都是16 进制的
// 	return fmt.Sprintf("ID: %s-Name: %s-Age : %d", e.Id, e.Name, e.Age) // %s 取string 占位。%d 取int
// }

// 这个和上门比较的好处就是 你获取的是 同一块内存地址 因为你是指针 ，指针实际上是指向同一片内存地址的 ,所以没有复制多余的存储空间 建议使用
func (e *Employee) String() string {
	// 代码 x处 下面我们来取它的内存地址
	fmt.Printf("Address is %x", unsafe.Pointer(&e.Name)) // 这个 unsafa包自己下去 %x是去16进制的数据，一般来说内存地址都是16 进制的

	return fmt.Sprintf("ID: %s-Name: %s-Age : %d", e.Id, e.Name, e.Age) // %s 取string 占位。%d 取int
}

func TestFuncMethod(t *testing.T) {
	e := Employee{"0", "Employ", 20}
	fmt.Printf("Address is %x", unsafe.Pointer(&e.Name)) // 这个 unsafa包自己下去 %x是去16进制的数据，一般来说内存地址都是16 进制的

	t.Log(e.String())
}
