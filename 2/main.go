package main

import "fmt"

func main() {
	people := map[string]int{
		"Иван":  20,
		"Мария": 25,
		"Петр":  18,
	}
	people["Александр"] = 28
	fmt.Println("Средний возраст: ", AverageAge(people))
}

func AverageAge(ArgMap map[string]int) float64 {
	answer := 0.0
	count := 0
	for _, age := range ArgMap {
		answer += float64(age)
		count++
	}
	return answer / float64(count)
}
