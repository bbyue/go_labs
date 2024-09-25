package main

import "fmt"

func main() {
	people := map[string]int{
		"Иван":  20,
		"Мария": 25,
		"Петр":  18,
	}

	people["Александр"] = 28

	fmt.Println("Все записи:")
	for name, age := range people {
		fmt.Printf("%s: %d\n", name, age)
	}
}
