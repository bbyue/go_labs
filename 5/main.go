package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите несколько чисел:")
	input, _ := reader.ReadString('\n')
	numbers := strings.Fields(strings.TrimSpace(input))
	sum := 0
	for _, number := range numbers {
		num, err := strconv.Atoi(number)
		if err != nil {
			fmt.Println("Неверный ввод:", err)
			return
		}
		sum += num
	}

	fmt.Println("Сумма чисел:", sum)
}
