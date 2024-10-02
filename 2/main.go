package main

import (
	"fmt"
	"time"
)

func Fibonacci(c chan int) {
	x, y := 0, 1
	for i := 0; i < 10; i++ {
		c <- x
		x = y
		y = x + y
	}
	close(c)
}
func Print(c chan int) {
	for num := range c {
		fmt.Println(num)
	}
}
func main() {
	c := make(chan int)
	go Fibonacci(c)
	go Print(c)
	time.Sleep(1 * time.Second)
}
