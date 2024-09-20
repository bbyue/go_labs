package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("1. Current time is: ", now)
	var myInt int = 32
	int := 64
	var myFloat64 float64 = 5.03485
	float64 := 4.59238483
	var myString string = "Hello"
	string := "World"
	var myBool bool = true
	bool := false
	fmt.Println()
	fmt.Println("2, 3. int = ", myInt, int)
	fmt.Println("float64 = ", myFloat64, float64)
	fmt.Println("string = ", myString, string)
	fmt.Println("bool = ", myBool, bool)
	fmt.Println()
	arythmeticOperations(63, 5)
	floatOperations(4.308, 5.394)
	average(myFloat64, float64, 9.456576)
}

func arythmeticOperations(a int, b int) {
	fmt.Println("4. Arythmetic operations: a = ", a, "b = ", b)
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println()
}

func floatOperations(a float64, b float64) {
	fmt.Println("5. Float operations: a = ", a, "b = ", b)
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println()
}

func average(a float64, b float64, c float64) {
	fmt.Println("6. Average: a = ", a, "b = ", b, "c = ", c)
	fmt.Println(a + b + c/3)
	fmt.Println()
}
