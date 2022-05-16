package test

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Space() {
	fmt.Println("-----")
}

func (p *Pet) SpaceTO(host string) {
	p.Space()
	fmt.Println("hello哇", host)
}

// 定义子类
// type Dog struct {
// 	p *Pet
// }

// func (d *Dog) Space() {
// 	d.p.Space()
// }

// func (d *Dog) SpaceTO(host string) {
// 	d.Space()
// 	fmt.Println("-----子", host)
// }

// 我们有更加简单的方式实现 组合
type Dog struct {
	Pet // 匿名嵌套类型 类似js中的原型
}

// 如果需要覆 写 就需要自己重新定义
func (d *Dog) Space() {
	fmt.Println("子复写-----")
}

func (d *Dog) SpaceTO(host string) {
	d.Space()
	fmt.Println("-----子", host)
}

func TestWrapDog(t *testing.T) {
	dog := new(Dog)
	dog.SpaceTO("Chan")
}
