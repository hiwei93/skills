package main

// 练习：错误(https://tour.go-zh.org/methods/9)

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := x / 2
	i := 0
	stop := false
	for !stop && i < 500 {
		temp := z
		z -= (z*z - x) / (2 * z)
		if abs(temp-z) < 0.0000000000001 {
			stop = true
		}
		i++
		//fmt.Printf("times: %v; z: %v\n", i, z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
