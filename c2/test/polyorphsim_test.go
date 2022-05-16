package test

import (
	"fmt"
	"testing"
)

type Code string

type ProgramerX interface {
	writeCode() Code
}

// 我们有一个程序要复合的参数p 得复合这个接口
// 不同的人实现它，就有不同的表现出来，这就是多态
func someCode(p ProgramerX) {
	fmt.Printf("%T %v\n", p, p.writeCode())
}

type GoProgramerX struct{}
type JavaProgramerX struct{}

func (goP *GoProgramerX) writeCode() Code {
	return "fmt.Print()"
}

func (javaP *JavaProgramerX) writeCode() Code {
	return "System.out.Print()"
}

// 进入测试程序
func TestPos(t *testing.T) {
	goProg := new(GoProgramerX)
	// 注意啊⚠️ 前面我们 说过 new( struct ) 和 struct{} 区别
	// 前置返回的是指针，后者返回的是 实例

	// 如果你  struct{  } ,那么就得这样写
	// goProg := &GoProgramerX{}

	javaProg := new(JavaProgramerX)
	someCode(goProg)
	someCode(javaProg)
}

//空接口？空接口的作用一般都是 为了断言 和转 any
func doSomething(p interface{}) {
	// if i, ok := p.(int); ok {   // p .(type) 是断言 你可以用swtich 也可以使用 if
	// 	fmt.Println("integer",i)
	// }
	switch v := p.(type) {
	case int:
		fmt.Println("Interger", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("unkouwn")
	}
}

func TestEmptyInterface(t *testing.T) {
	doSomething(10)
	doSomething("10")
}
