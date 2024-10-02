package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Factorial(num int) {
	i := 1
	product := 1
	for i < num+1 {
		product *= i
		i++
	}
	fmt.Println("Факториал числа", num, "=", product)
}
func Random(num int) {
	fmt.Println("Случайное число:", rand.Intn(num))
}
func ArythmeticProgression(a0 int, an int, n int) {
	sum := ((float64(a0+an) / 2) * float64(n))
	fmt.Println("Сумма арифметической прогрессии:", sum)
}
func main() {
	go Factorial(7)
	go Random(100)
	go ArythmeticProgression(0, 16, 9)
	time.Sleep(1 * time.Second)
}
