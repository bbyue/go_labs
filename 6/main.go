package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	strSlice1 := []string{"Написать", "программу", "которая", "создает", "срез", "из", "строк", "и", "находит", "самую", "длинную", "строку"}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите несколько слов:")
	input, _ := reader.ReadString('\n')
	strSlice2 := strings.Fields(strings.TrimSpace(input))
	fmt.Println("Самое длинное слово:", findLongest(strSlice1))
	fmt.Println("Самое длинное слово:", findLongest(strSlice2))
}

func findLongest(str []string) string {
	answer := ""
	for i := 0; i < len(str); i++ {
		if len(str[i]) >= len(answer) {
			answer = str[i]
		}
	}
	return answer
}
