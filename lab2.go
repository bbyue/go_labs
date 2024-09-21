package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rectangle struct {
	width  int
	height int
}

func (r Rectangle) rectangleArea() int {
	return r.width * r.height
}

func main() {
	isEven()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println()
	fmt.Println("2. Введите число: ")
	input, _ := reader.ReadBytes('\n')
	inputStr := strings.TrimSpace(string(input))
	num, err := strconv.Atoi(inputStr)
	if err != nil {
		fmt.Println("Ошибка преобразования:", err)
	}
	fmt.Println(whatNumber(num))
	fmt.Println()
	for10()
	fmt.Println()
	fmt.Println("4. Длина строки = ", strLength("Hello, World!"))
	fmt.Println()
	rect := Rectangle{width: 5, height: 4}
	fmt.Println("5. Площадь треугольника = ", rect.rectangleArea())
	fmt.Println()
	fmt.Println("Среднее значение = ", average())
}
func isEven() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("1. Введите число: ")
	input, _ := reader.ReadBytes('\n')
	inputStr := strings.TrimSpace(string(input))
	num, err := strconv.Atoi(inputStr)
	if err != nil {
		fmt.Println("Ошибка преобразования:", err)
	}
	if num&2 == 0 {
		fmt.Println("Число четное")
		return true
	} else {
		fmt.Println("Число нечетное")
		return false
	}

}
func whatNumber(a int) string {
	if a > 0 {
		return "Positive"
	} else if a < 0 {
		return "Negative"
	} else {
		return "Zero"
	}
}

func for10() {
	fmt.Println("3.")
	i := 1
	for i < 11 {
		fmt.Println(i)
		i++
	}
}

func strLength(str string) int {
	return len(str)
}
func average() float64 {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("6. Введите число a: ")
	input, _ := reader.ReadBytes('\n')
	inputStr := strings.TrimSpace(string(input))
	num1, err := strconv.Atoi(inputStr)
	if err != nil {
		fmt.Println("Ошибка преобразования:", err)
	}
	fmt.Println("Введите число b: ")
	input2, _ := reader.ReadBytes('\n')
	inputStr2 := strings.TrimSpace(string(input2))
	num2, err := strconv.Atoi(inputStr2)
	if err != nil {
		fmt.Println("Ошибка преобразования:", err)
	}
	ans := float64(num1) + float64(num2)
	var float float64 = ans / 2
	return float
}
