package main

import (
	"fmt"
	mathutils "mathutils/package"
)

func main() {
	var num int
	fmt.Println("Введите число: ")
	fmt.Scan(&num)
	fmt.Println("Факториал числа ", num, ": ", mathutils.Factorial(num))
}
