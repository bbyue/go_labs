package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func (r Person) PrintPerson() {
	fmt.Println("Имя:", r.name, " Возраст:", r.age)
}
func main() {
	ivan := Person{"Иван", 20}
	ivan.PrintPerson()
}
