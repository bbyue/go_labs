package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type task struct {
	line string
}

func Reverse(str string) string {
	ans := make([]byte, len(str))
	i := 0
	j := len(str) - 1

	for i < len(str) {
		ans[i] = str[j]
		i++
		j--
	}
	return string(ans)
}

func worker(tasks <-chan task) {
	for task := range tasks {
		reversed := Reverse(task.line)
		fmt.Println(reversed)
	}
}

func main() {
	var num int
	fmt.Print("Введите количество воркеров: ")
	fmt.Scan(&num)

	tasks := make(chan task)

	for i := 0; i < num; i++ {
		go worker(tasks)
	}

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tasks <- task{line: scanner.Text()}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения файла:", err)
	}

	file.Close()
	close(tasks)
	time.Sleep(1 * time.Second)
}
