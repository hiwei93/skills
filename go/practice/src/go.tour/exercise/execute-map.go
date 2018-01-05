package main

// 练习：map(https://tour.go-zh.org/moretypes/20)

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	fMap := make(map[string]int)
	fields := strings.Fields(s)
	for _, field := range fields {
		fMap[field] = fMap[field] + 1
	}
	return fMap
}

func main() {
	wc.Test(WordCount)
}
