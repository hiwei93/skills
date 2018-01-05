package main

// 练习：斐波纳契闭包(https://tour.go-zh.org/moretypes/23)

import "fmt"

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {
	a := 0
	b := 1
	return func() int {
		temp := a
		a = b
		b = temp + b
		return temp
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
