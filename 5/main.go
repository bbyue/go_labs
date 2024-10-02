package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(float64(c.radius), 2)
}
func printShapeInfo(s []Shape) {
	for i := 0; i < len(s); i++ {
		fmt.Println("Площадь: ", s[i].Area())
	}
}
func main() {
	shapes := make([]Shape, 3)
	rect1 := Rectangle{3, 4}
	circle1 := Circle{8}
	rect2 := Rectangle{10, 15}
	shapes[0] = rect1
	shapes[1] = circle1
	shapes[2] = rect2
	printShapeInfo(shapes)
}
