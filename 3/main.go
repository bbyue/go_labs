package main

import "fmt"

func main() {
	people := map[string]int{
		"Иван":  20,
		"Мария": 25,
		"Петр":  18,
	}
	people["Александр"] = 28
	fmt.Println("До удаления:")
	for name, age := range people {
		fmt.Printf("%s: %d\n", name, age)
	}
	fmt.Println("После удаления:")
	delete(people, "Иван")
	for name, age := range people {
		fmt.Printf("%s: %d\n", name, age)
	}
}
