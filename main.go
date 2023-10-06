package main

import (
	"fmt"
	"strconv"
)

func swapper(a, b string) (int, int, bool) {
	r, err := strconv.Atoi(a)
	if err != nil {
		return 0, 0, true
	}
	r2, err := strconv.Atoi(b)
	if err != nil {
		return 0, 0, true
	}
	return r, r2, false
}

func main() {
	a, b, err := swapper("123", "333")
	if err {
		fmt.Println("asdasd")
		return
	}
	fmt.Println(a+1, b+1000)
	return
}
