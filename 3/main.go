package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (r Circle) CircleArea() float64 {
	return math.Pi * math.Pow(r.radius, 2)
}

func main() {
	circle := Circle{radius: 5}
	fmt.Println("Площадь круга:", circle.CircleArea())
}
