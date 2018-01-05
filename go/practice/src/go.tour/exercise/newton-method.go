package main

// 练习：循环和函数

import (
	"fmt"
)

// Sqrt 初步实现了牛顿法求平方根
func Sqrt(x float64) float64 {
	z := 1.0
	i := 0
	for i < 10 {
		z -= (z*z - x) / (2 * z)
		i++
		fmt.Printf("times: %v; z: %v\n", i, z)
	}
	return z
}

// abs 是求一个float64类型的绝对值，帮助SqrtAdvance的一个函数
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

// SqrtAdvance 进一步实现了牛顿法求平方根，添加了精度值参数precision
func SqrtAdvance(x float64, precision float64) float64 {
	z := x / 2
	i := 0
	stop := false
	for !stop && i < 500 {
		temp := z
		z -= (z*z - x) / (2 * z)
		if abs(temp-z) < precision {
			stop = true
		}
		i++
		fmt.Printf("times: %v; z: %v\n", i, z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(SqrtAdvance(3, 0.00000000001))
}
