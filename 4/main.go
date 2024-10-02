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
func printShapeInfo(s Shape) {
	fmt.Println("Площадь:", s.Area())
}
func main() {
	rect := Rectangle{width: 5, height: 6}
	circle := Circle{radius: 5}
	printShapeInfo(rect)
	printShapeInfo(circle)
}
