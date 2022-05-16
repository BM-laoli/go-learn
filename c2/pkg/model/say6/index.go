package sya6

import "fmt"

func SayHello() {
	fmt.Println("你好-public")
}

func sayHello1() {
	fmt.Println("你好-private")
	sayHello1()
}
