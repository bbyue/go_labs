package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 5)
	for i := 0; i < 5; i++ {
		slice[i] = i + 1
	}
	slice = append(slice, 6) //добавление элемента
	fmt.Println("До удаления: ", slice)
	fmt.Println()
	slice = append(slice[:2], slice[3:]...)
	fmt.Println("После удаления: ", slice)
}
