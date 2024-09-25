package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var array [5]int
	for i := 0; i < 5; i++ {
		array[i] = rand.Intn(100)
		fmt.Println(array[i])
	}
}
