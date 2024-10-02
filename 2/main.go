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
func (r *Person) Birthday() {
	r.age++
}
func main() {
	ivan := Person{"Иван", 20}
	fmt.Println("До дня рождения:")
	ivan.PrintPerson()
	fmt.Println("После дня рождения:")
	ivan.Birthday()
	ivan.PrintPerson()
}
