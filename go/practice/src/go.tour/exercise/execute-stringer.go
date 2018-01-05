package main

// 练习：Stringers(https://tour.go-zh.org/methods/7)

import "fmt"

type IPAddr [4]byte

func (ip IPAddr) String() string {
	var s string
	for i, value := range ip {
		if i != 0 {
			s = s + "."
		}

		s = fmt.Sprintf("%s%v", s, value)
	}
	return s
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
