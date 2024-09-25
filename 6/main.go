package main

import (
	"fmt"
)

func main() {
	var num [10]int
	for i := 0; i < len(num); i++ {
		num[i] = i
	}
	Backwards(num[:])
}

func Backwards(num []int) {
	for i := len(num) - 1; i > -1; i-- {
		fmt.Println(num[i])
	}
}
